package ingestion

import (
	"net/http"
	"strconv"
	"time"

	"github.com/iris-analytics/iris-backend/internal/domain/repository"
	"github.com/iris-analytics/iris-backend/internal/infrastructure/transformer"

	"github.com/labstack/echo"
)

var transparentPixel = []byte("\x47\x49\x46\x38\x39\x61\x01\x00\x01\x00\x80\x00\x00\x00\x00\x00\x00\x00\x00\x21\xF9\x04\x01\x00\x00\x00\x00\x2C\x00\x00\x00\x00\x01\x00\x01\x00\x00\x02\x02\x44\x01\x00\x3B")

// EventHandler handles the request to store an event
type EventHandler struct {
	repository repository.EventRepositoryInterface
}

// NewEventHandler creates a new Handler
func NewEventHandler(r repository.EventRepositoryInterface) *EventHandler {
	h := &EventHandler{repository: r}
	return h
}

// Ingest Transforms a request into something that can be saved
func (h EventHandler) Ingest(c echo.Context) error {

	cd, _ := strconv.ParseInt(c.QueryParam("cd"), 10, 16)
	md, _ := strconv.ParseBool(c.QueryParam("md"))
	tz, _ := strconv.ParseInt(c.QueryParam("tz"), 10, 16)

	req := transformer.Request{
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

	r := req.MakeEvent()

	go func() {
		_, err := h.repository.Persist(r)
		if err != nil {
			c.Logger().Error(err)
		}
	}()

	return c.Blob(http.StatusOK, "image/gif", transparentPixel)
}
