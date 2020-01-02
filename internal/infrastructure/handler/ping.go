package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// Ping handles health check
type Ping struct {
}

// HandlePing does pong
func (h Ping) HandlePing(c echo.Context) error {

	return c.JSON(http.StatusOK, "PONG")
}
