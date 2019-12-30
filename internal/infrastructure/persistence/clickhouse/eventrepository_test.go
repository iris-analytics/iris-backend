package clickhouse

import (
	"testing"

	"github.com/iris-analytics/iris-backend/internal/domain/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewEventRepositoryReturnsProperInstance(t *testing.T) {
	// Arrange
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Act
	r := NewEventRepository(db)
	// Assert
	assert.Implements(t, (*repository.EventRepositoryInterface)(nil), r)
}
