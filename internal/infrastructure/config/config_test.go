package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() {

}
func TestMain(m *testing.M) {

}

func TestNewConfig(t *testing.T) {

	// Arrange
	os.Setenv("IRIS_CLICKHOUSE_DSN", "foo")
	os.Setenv("IRIS_LISTEN_BINDING", "bar")
	// Act
	config, _ := NewConfig()
	// Assert
	assert.Equal(t, "foo", config.ClickHouseDSN)
	assert.Equal(t, "bar", config.ListenBinding)
}
