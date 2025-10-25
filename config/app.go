package config

import (
	"github.com/easy-attend-serviceV3/app/modules/example"
	exampletwo "github.com/easy-attend-serviceV3/app/modules/example-two"
	"github.com/easy-attend-serviceV3/internal/log"
	"github.com/easy-attend-serviceV3/internal/otel/collector"
)

// Config is a struct that contains all the configuration of the application.
type Config struct {
	Database Database

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
