package config

import (
	"errors"
	"os"
)

// Config stores app configuration
type Config struct {
	ClickHouseDSN string
	ListenBinding string
}

// NewConfig returns a new Config struct
func NewConfig() (*Config, error) {

	clickHouseDSN, clickHouseDSNExists := os.LookupEnv("IRIS_CLICKHOUSE_DSN")
	listenBinding, listenBindingExists := os.LookupEnv("IRIS_LISTEN_BINDING")

	if !clickHouseDSNExists {
		return nil, errors.New("Environment variable IRIS_CLICKHOUSE_DSN is not set")
	}
	if !listenBindingExists {
		return nil, errors.New("Environment variable IRIS_LISTEN_BINDING is not set")
	}

	return &Config{
		ClickHouseDSN: clickHouseDSN,
		ListenBinding: listenBinding,
	}, nil
}
