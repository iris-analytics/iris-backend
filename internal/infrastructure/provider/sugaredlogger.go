package provider

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// GetSugaredLogger ...
func GetSugaredLogger(appName string) *zap.SugaredLogger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "application",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "context",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
	// mutex on logging
	stdoutSync := zapcore.Lock(os.Stdout)
	stdErrSync := zapcore.Lock(os.Stderr)
	// change log threshold based on app environment
	stdoutEnabler := func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel && lvl >= zapcore.InfoLevel
	}
	stderrEnabler := func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	}
	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, stdoutSync, zap.LevelEnablerFunc(stdoutEnabler)),
		zapcore.NewCore(jsonEncoder, stdErrSync, zap.LevelEnablerFunc(stderrEnabler)),
	)
	logger := zap.New(core)
	logger = logger.Named(appName)
	sugaredLogger := logger.Sugar()

	hostname, _ := os.Hostname()

	sugaredLogger = sugaredLogger.With(
		"host", hostname,
		"type", "log",
	)
	sugaredLogger.Debug("Logging active")
	return sugaredLogger
}
