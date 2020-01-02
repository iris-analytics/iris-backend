package provider

import (
	"go.uber.org/zap"
)

// GetSugaredLogger returns an instance of zap's sugared logger
func GetSugaredLogger(appName string) *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	return sugar
}
