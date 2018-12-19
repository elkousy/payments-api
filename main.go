package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"runtime/pprof"
	"strconv"
	"syscall"
	"time"

	"github.com/elkousy/payments-api/payments"
	"github.com/elkousy/payments-api/utility/config"
	"github.com/elkousy/payments-api/utility/logger"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	errc := make(chan error)
	var err error

	// handle syscall signals for a gracefull shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// init database
	db, err := payments.DbConnect()
	if err != nil {
		logger.LogStdErr.Error(errors.Wrap(err, "error when connecting to postgres"))
		os.Exit(0)
	}

	// init repository
	repository := payments.NewPaymentRepository(db)
	defer payments.DbClose(db)
	payments.DbMigrate(db)

	// init service
	svc, err := payments.NewPaymentService(repository)
	if err != nil {
		errc <- err
	}

	// build api endpoints
	endpoints := payments.MakeEndpoints(svc)

	// Instances a new HTTP server
	go func() {
		httpAddr := ":" + strconv.Itoa(config.AppPort)
		mux := mux.NewRouter()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Welcome to the Payments API!")
		})

		// init and register to the router the various endpoints
		payments.MakeHTTPHandler(endpoints, mux)

		logger.LogStdOut.Info(fmt.Sprintf("The %s has started on port %s", config.AppName, httpAddr))

		s := &http.Server{
			Addr:         httpAddr,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			Handler:      mux,
		}

		errc <- s.ListenAndServe()
	}()

	// launch a dedicated webserver for observability , i.e. health check and metrics
	go func() {
		opsHTTPAddr := ":" + strconv.Itoa(config.OpsPort)
		mux := mux.NewRouter()

		mux.Handle("/metrics", promhttp.Handler())
		mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		logger.LogStdOut.Info(fmt.Sprintf("The ops server has started on port %s", opsHTTPAddr))
		errc <- http.ListenAndServe(opsHTTPAddr, mux)
	}()

	
	// launch a dedicated webserver for pprof
	go func() {
		debugHTTPAddr := ":" + strconv.Itoa(config.DebugPort)
		logger.LogStdOut.Info(fmt.Sprintf("The pprof server will start on port  %s", debugHTTPAddr))
		http.HandleFunc("/_stack", getStackTraceHandler)
		http.ListenAndServe(debugHTTPAddr, nil)
	}()

	fmt.Println("exit", <-errc)
}

func getStackTraceHandler(w http.ResponseWriter, r *http.Request) {
	stack := debug.Stack()
	w.Write(stack)
	pprof.Lookup("goroutine").WriteTo(w, 2)
}
