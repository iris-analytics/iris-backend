package handler

import (
	"net/http"

	"github.com/iris-analytics/iris-backend/internal/application/usecase"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

var transparentPixel = []byte("\x47\x49\x46\x38\x39\x61\x01\x00\x01\x00\x80\x00\x00\x00\x00\x00\x00\x00\x00\x21\xF9\x04\x01\x00\x00\x00\x00\x2C\x00\x00\x00\x00\x01\x00\x01\x00\x00\x02\x02\x44\x01\x00\x3B")

// RecordEvent ...
type RecordEvent struct {
	UseCase usecase.RecordEventInterface
	Logger  *zap.SugaredLogger
}

// HandleRecordEvent ...
func (handler *RecordEvent) HandleRecordEvent(c echo.Context) error {

	err := handler.UseCase.Execute(c)
	if err != nil {
		handler.Logger.Error(err)
	}

	return c.Blob(http.StatusOK, "image/gif", transparentPixel)
}
