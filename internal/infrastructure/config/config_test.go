package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func cleanEnvVars() {
	os.Unsetenv("IRIS_CLICKHOUSE_DSN")
	os.Unsetenv("IRIS_CLICKHOUSE_TABLE")
	os.Unsetenv("IRIS_LISTEN_BINDING")
	os.Unsetenv("IRIS_RECORDER_PATH")
}

func (suite *ConfigTestSuite) SetupTest() {
	cleanEnvVars()
}
func (suite *ConfigTestSuite) TearDownTest() {
	cleanEnvVars()
}
func (suite *ConfigTestSuite) TearDownSuite() {
	cleanEnvVars()
}

func (suite *ConfigTestSuite) TestConfigPicksEnvVarsCorrectly() {
	// Arrange
	os.Setenv("IRIS_CLICKHOUSE_DSN", "a")
	os.Setenv("IRIS_CLICKHOUSE_TABLE", "b")
	os.Setenv("IRIS_LISTEN_BINDING", "c")
	os.Setenv("IRIS_RECORDER_PATH", "d")
	// Act
	config, _ := NewConfig()
	// Assert
	suite.Equal("a", config.ClickHouseDSN)
	suite.Equal("b", config.ClickHouseTable)
	suite.Equal("c", config.ListenBinding)
	suite.Equal("d", config.RecorderPath)
}

func (suite *ConfigTestSuite) TestErrorIfClickHouseDSNVarIsNotSet() {
	// Arrange
	os.Setenv("IRIS_CLICKHOUSE_TABLE", "b")
	os.Setenv("IRIS_LISTEN_BINDING", "c")
	os.Setenv("IRIS_RECORDER_PATH", "d")
	// Act
	_, err := NewConfig()
	// Assert
	suite.NotNil(err)
	suite.Equal("Environment variable IRIS_CLICKHOUSE_DSN is not set", err.Error())
}
func (suite *ConfigTestSuite) TestErrorIfClickHouseTableVarIsNotSet() {
	// Arrange
	os.Setenv("IRIS_CLICKHOUSE_DSN", "a")
	os.Setenv("IRIS_LISTEN_BINDING", "c")
	os.Setenv("IRIS_RECORDER_PATH", "d")
	// Act
	_, err := NewConfig()
	// Assert
	suite.NotNil(err)
	suite.Equal("Environment variable IRIS_CLICKHOUSE_TABLE is not set", err.Error())
}
func (suite *ConfigTestSuite) TestErrorIfListenBindingVarIsNotSet() {
	// Arrange
	os.Setenv("IRIS_CLICKHOUSE_DSN", "a")
	os.Setenv("IRIS_CLICKHOUSE_TABLE", "b")
	os.Setenv("IRIS_RECORDER_PATH", "d")
	// Act
	_, err := NewConfig()
	// Assert
	suite.NotNil(err)
	suite.Equal("Environment variable IRIS_LISTEN_BINDING is not set", err.Error())
}
func (suite *ConfigTestSuite) TestErrorIfRecorderPathVarIsNotSet() {
	// Arrange
	os.Setenv("IRIS_CLICKHOUSE_DSN", "a")
	os.Setenv("IRIS_CLICKHOUSE_TABLE", "b")
	os.Setenv("IRIS_LISTEN_BINDING", "c")
	// Act
	_, err := NewConfig()
	// Assert
	suite.NotNil(err)
	suite.Equal("Environment variable IRIS_RECORDER_PATH is not set", err.Error())
}

func TestConfigSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(ConfigTestSuite))
}
