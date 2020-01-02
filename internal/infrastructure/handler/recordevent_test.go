package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type DummyUseCase struct {
}

func (duc *DummyUseCase) Execute(c echo.Context) error {
	return nil
}

func getInvisiblePixel() string {
	return "GIF89a\x01\x00\x01\x00\x80\x00\x00\x00\x00\x00\x00\x00\x00!\xf9\x04\x01\x00\x00\x00\x00,\x00\x00\x00\x00\x01\x00\x01\x00\x00\x02\x02D\x01\x00;"
}

func TestRecordEventHandleRecordEventReturnsProperReponse(t *testing.T) {

	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/any", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := RecordEvent{UseCase: &DummyUseCase{}}
	// Act
	// Assert
	if assert.NoError(t, h.HandleRecordEvent(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, getInvisiblePixel(), rec.Body.String())
	}
}
