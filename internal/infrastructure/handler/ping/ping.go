package ping

import (
	"net/http"

	"github.com/labstack/echo"
)

// Handler handles health check
type Handler struct {
}

// Ping does pong
func (h Handler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "PONG")
}
