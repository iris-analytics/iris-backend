package repository

import "github.com/iris-analytics/iris-backend/internal/domain/entity"

// EventRepositoryInterface If a repo you want to be, this interface you must implement
type EventRepositoryInterface interface {
	Record(e *entity.Event) (*entity.Event, error)
}
