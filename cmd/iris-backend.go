package main

import (
	"github.com/iris-analytics/iris-backend/internal/application/usecase"
	"github.com/iris-analytics/iris-backend/internal/infrastructure/config"
	"github.com/iris-analytics/iris-backend/internal/infrastructure/handler"
	"github.com/iris-analytics/iris-backend/internal/infrastructure/persistence/clickhouse"
	"github.com/iris-analytics/iris-backend/internal/infrastructure/provider"

	"github.com/labstack/echo"
)

const appName = "iris-backend"

func main() {

	e := echo.New()
	e.HidePort = true
	logger := provider.GetSugaredLogger(appName)

	defer logger.Sync()

	e.HideBanner = true
	config, err := config.NewConfig()

	if err != nil {
		logger.Fatal(err)
	}

	// Record event
	httpClient := provider.GetPesterHTTPClient()
	eventRepo := clickhouse.NewEventRepository(
		httpClient,
		config.ClickHouseDSN,
		config.ClickHouseTable,
		logger,
	)
	recordEventHandler := handler.RecordEvent{
		UseCase: &usecase.RecordEvent{
			EventRepository: eventRepo,
			Logger:          logger,
		},
		Logger: logger,
	}

	// Record event happens in both GET and POST since it can be via tracking pixel (GET)
	// or document.beacon (navigator.sendBeacon). In both cases, data is sent in query string
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
