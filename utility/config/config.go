package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	logger "github.com/elkousy/payments-api/utility/logger"
	"github.com/spf13/viper"
)

var (
	AppName    string
	AppPort    int
	OpsPort    int
	DebugPort  int
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
	DBTimeout  int
)

func init() {
	InitConfig()
}

// InitConfig initializes all config vars
func InitConfig() {
	viper.SetDefault("AppName", "payments-api")
	viper.SetDefault("APP_PORT", 8080)
	viper.SetDefault("OPS_PORT", 8081)
	viper.SetDefault("DEBUG_PORT", 8082)

	var isDev bool
	switch strings.ToLower(os.Getenv("ENVIRONMENT")) {
	case "dev":
		isDev = true
	}
	if isDev {
		_, dirname, _, _ := runtime.Caller(0)
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		//viper.AddConfigPath(".")
		viper.AddConfigPath(filepath.Dir(dirname))
		err := viper.ReadInConfig()
		if err != nil {
			logger.LogStdErr.Error(err)
		}
	} else {
		viper.AutomaticEnv()
	}

	AppName = viper.GetString("AppName")
	AppPort = viper.GetInt("APP_PORT")
	OpsPort = viper.GetInt("OPS_PORT")
	DebugPort = viper.GetInt("DEBUG_PORT")

	// db configuration
	DBHost = viper.GetString("DB_HOST")
	DBPort = viper.GetInt("DB_PORT")
	DBName = viper.GetString("DB_NAME")
	DBUser = viper.GetString("DB_USER")
	DBPassword = viper.GetString("DB_PASSWORD")
	DBTimeout = viper.GetInt("DB_TIMEOUT")
}
