package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {

	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/ping", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := Handler{}
	// Act
	// Assert
	if assert.NoError(t, h.Ping(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"PONG\"\n", rec.Body.String())
	}
}
