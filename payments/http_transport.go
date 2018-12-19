package payments

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	apierrors "github.com/elkousy/payments-api/utility/errors"
	"github.com/elkousy/payments-api/utility/instrumenting"
)

const componentName = "payments"

// MakeHTTPHandler returns all http handler for the payments service
func MakeHTTPHandler(endpoints Endpoints, router *mux.Router) http.Handler {

	options := []kithttp.ServerOption{
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
		kithttp.ServerErrorEncoder(apierrors.LoggingErrorEncoder),
	}

	getPaymentHandler := instrumenting.Middleware(componentName, "get_payment_by_id", kithttp.NewServer(
		endpoints.GetPayment,
		decodeGetPaymentRequest,
		encodeOKResponse,
		options...,
	))

	getListOfPaymentsHandler := instrumenting.Middleware(componentName, "get_list_of_payments", kithttp.NewServer(
		endpoints.GetListOfPayments,
		decodeGetListOfPaymentsRequest,
		encodeOKResponse,
		options...,
	))

	updatePaymentHandler := instrumenting.Middleware(componentName, "put_payment", kithttp.NewServer(
		endpoints.UpdatePayment,
		decodeUpdatePaymentRequest,
		encodeAcceptedResponse,
		options...,
	))

	postPaymentHandler := instrumenting.Middleware(componentName, "post_payment", kithttp.NewServer(
		endpoints.PostPayment,
		decodePostPaymentRequest,
		encodeCreatedResponse,
		options...,
	))

	deletePaymentHandler := instrumenting.Middleware(componentName, "delete_payment", kithttp.NewServer(
		endpoints.DeletePayment,
		decodeDeletePaymentRequest,
		encodeAcceptedResponse,
		options...,
	))

	r := router.PathPrefix("/v1/payments").Subrouter().StrictSlash(true)
	{
		r.Handle("/{id}/", getPaymentHandler).Methods(http.MethodGet)
		r.Handle("/", getListOfPaymentsHandler).Methods(http.MethodGet)
		r.Handle("/", postPaymentHandler).Methods(http.MethodPost)
		r.Handle("/{id}/", updatePaymentHandler).Methods(http.MethodPut)
		r.Handle("/{id}/", deletePaymentHandler).Methods(http.MethodDelete)
	}

	return r
}

func decodeGetPaymentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	return GetPaymentRequest{PaymentID: id}, nil
}

func decodeGetListOfPaymentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := GetListOfPaymentsRequest{}
	return req, nil
}

func decodePostPaymentRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrInvalidBody
	}
	return req, nil
}

func decodeUpdatePaymentRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req UpdatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrInvalidBody
	}
	vars := mux.Vars(r)
	req.PaymentID = vars["id"]
	return req, nil
}

func decodeDeletePaymentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	return DeletePaymentRequest{PaymentID: id}, nil
}

func encodeOKResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func encodeAcceptedResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusAccepted)
	return nil
}

func encodeCreatedResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}
