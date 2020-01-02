package usecase

import "github.com/labstack/echo"

// RecordEventInterface is the simple interface for all usecases
type RecordEventInterface interface {
	Execute(c echo.Context) error
}
