package usecase

import (
	"strconv"
	"time"

	"github.com/iris-analytics/iris-backend/internal/domain/repository"
	"github.com/iris-analytics/iris-backend/internal/infrastructure/transformer"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// RecordEvent will record the event sent to the backend
type RecordEvent struct {
	EventRepository repository.EventRepositoryInterface
	Logger          *zap.SugaredLogger
}

// Execute will execute the use case
func (useCase *RecordEvent) Execute(c echo.Context) error {

	cd, _ := strconv.ParseInt(c.QueryParam("cd"), 10, 16)
	md, _ := strconv.ParseBool(c.QueryParam("md"))
	tz, _ := strconv.ParseInt(c.QueryParam("tz"), 10, 16)

	eventRequest := transformer.Request{
		AccountID:        c.QueryParam("id"),
		VisitorID:        c.QueryParam("uid"),
		SessionID:        c.QueryParam("sid"),
		EventType:        c.QueryParam("ev"),
		EventData:        c.QueryParam("ed"),
		DocumentLocation: c.QueryParam("dl"),
		ReferrerLocation: c.QueryParam("rl"),
		DocumentEncoding: c.QueryParam("de"),
		ScreenResolution: c.QueryParam("sr"),
		ViewPort:         c.QueryParam("vp"),
		ColorDepth:       uint16(cd),
		DocumentTitle:    c.QueryParam("dt"),
		BrowserName:      c.QueryParam("bn"),
		IsMobileDevice:   md,
		UserAgent:        c.QueryParam("ua"),
		TimeZoneOffset:   int16(tz),
		Utm:              c.QueryParam("utm"),
		Timestamp:        time.Now(),
		IPAddress:        c.RealIP(),
	}

	event := eventRequest.MakeEvent()

	go func() {
		_, err := useCase.EventRepository.Record(event)
		if err != nil {
			useCase.Logger.Error(err)
		}
	}()

	return nil

}
