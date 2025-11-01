package config

import (
	"github.com/easy-attend-serviceV3/app/modules/example"
	exampletwo "github.com/easy-attend-serviceV3/app/modules/example-two"
	"github.com/easy-attend-serviceV3/internal/log"
	"github.com/easy-attend-serviceV3/internal/otel/collector"
)

// JWTConfig contains JWT-related configuration
type JWTConfig struct {
	SecretKey            string
	AccessTokenExpiry    int // in hours
	RefreshTokenExpiry   int // in hours
	Issuer              string
}

// Config is a struct that contains all the configuration of the application.
type Config struct {
	Database Database
	JWT      JWTConfig

	AppName      string
	AppKey       string
	AppEnv       string
	AppEnvPrefix string
	Debug        bool

	Port           int
	HttpJsonNaming string

	SslCaPath      string
	SslPrivatePath string
	SslCertPath    string

	Otel collector.Config

	// Kafka dto.Kafka
	Log log.Option

	Example example.Config

	ExampleTwo exampletwo.Config
}

var App = Config{
	Database: database,
	// Kafka:    kafka,
	JWT: JWTConfig{
		SecretKey:            "your-super-secret-jwt-key-change-this-in-production",
		AccessTokenExpiry:    24,   // 24 hours
		RefreshTokenExpiry:   168,  // 7 days (24 * 7)
		Issuer:              "easy-attend-service",
	},

	AppName: "go_app",
	Port:    8080,
	AppKey:  "secret",
	AppEnv:  "development",
	Debug:   false,

	HttpJsonNaming: "snake_case",

	SslCaPath:      "mcop/cert/ca.pem",
	SslPrivatePath: "mcop/cert/server.pem",
	SslCertPath:    "mcop/cert/server-key.pem",

	Otel: collector.Config{
		CollectorEndpoint: "",
		LogMode:           "noop",
		TraceMode:         "noop",
		MetricMode:        "noop",
		TraceRatio:        0.01,
	},
}
