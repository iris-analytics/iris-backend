package config

import (
	"errors"
	"os"
)

// Config stores app configuration
type Config struct {
	ClickHouseDSN   string
	ClickHouseTable string
	ListenBinding   string
	RecorderPath    string
}

// NewConfig returns a new Config struct
func NewConfig() (*Config, error) {

	clickHouseDSN, clickHouseDSNExists := os.LookupEnv("IRIS_CLICKHOUSE_DSN")
	clickHouseTable, clickhouseTableExists := os.LookupEnv("IRIS_CLICKHOUSE_TABLE")
	listenBinding, listenBindingExists := os.LookupEnv("IRIS_LISTEN_BINDING")
	recorderPath, recorderPathExists := os.LookupEnv("IRIS_RECORDER_PATH")

	if !clickHouseDSNExists {
		return nil, errors.New("Environment variable IRIS_CLICKHOUSE_DSN is not set")
	}
	if !clickhouseTableExists {
		return nil, errors.New("Environment variable IRIS_CLICKHOUSE_TABLE is not set")
	}
	if !listenBindingExists {
		return nil, errors.New("Environment variable IRIS_LISTEN_BINDING is not set")
	}
	if !recorderPathExists {
		return nil, errors.New("Environment variable IRIS_RECORDER_PATH is not set")
	}

	return &Config{
		ClickHouseDSN:   clickHouseDSN,
		ClickHouseTable: clickHouseTable,
		ListenBinding:   listenBinding,
		RecorderPath:    recorderPath,
	}, nil
}
