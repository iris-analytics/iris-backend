package main

import (
	"database/sql"
	"os"

	"github.com/iris-analytics/iris-backend/internal/application/usecase"
	"github.com/iris-analytics/iris-backend/internal/infrastructure/config"
	"github.com/iris-analytics/iris-backend/internal/infrastructure/handler"
	"github.com/iris-analytics/iris-backend/internal/infrastructure/persistence/clickhouse"

	"github.com/labstack/echo"
	_ "github.com/mailru/go-clickhouse"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const appName = "iris-backend"

func main() {

	e := echo.New()
	e.HidePort = true
	logger := getSugaredLogger()

	defer logger.Sync()

	e.HideBanner = true
	config, err := config.NewConfig()

	if err != nil {
		logger.Fatal(err)
	}

	db, err := sql.Open("clickhouse", config.ClickHouseDSN)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Record event
	eventRepo := clickhouse.NewEventRepository(db, config.ClickHouseTable)
	recordEventHandler := handler.RecordEvent{
		UseCase: &usecase.RecordEvent{EventRepository: eventRepo},
	}
	e.GET(config.RecorderPath, recordEventHandler.HandleRecordEvent).Name = "RecordeventGET"
	e.POST(config.RecorderPath, recordEventHandler.HandleRecordEvent).Name = "RecordeventPOST"

	// Ping
	pingHandler := handler.Ping{}
	e.GET("/ping", pingHandler.HandlePing).Name = "ping"

	logger.Info(appName + " running on port " + config.ListenBinding)

	err = e.Start(":" + config.ListenBinding)

	if err != nil {
		logger.Fatal(err)
	}

}

func getSugaredLogger() *zap.SugaredLogger {
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
