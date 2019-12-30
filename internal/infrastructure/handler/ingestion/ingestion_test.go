package ingestion

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iris-analytics/iris-backend/internal/domain/entity"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type RepoMockNoError struct{}

func (r *RepoMockNoError) Persist(e *entity.Event) (*entity.Event, error) {
	return e, nil
}

type RepoMockWithError struct{}

func (r *RepoMockWithError) Persist(e *entity.Event) (*entity.Event, error) {
	return nil, errors.New("Failed to persist")
}

func getInvisiblePixel() string {
	return "GIF89a\x01\x00\x01\x00\x80\x00\x00\x00\x00\x00\x00\x00\x00!\xf9\x04\x01\x00\x00\x00\x00,\x00\x00\x00\x00\x01\x00\x01\x00\x00\x02\x02D\x01\x00;"
}

func TestIngestReturnsProperReponseIfFailsToPersist(t *testing.T) {

	// Arrange
	e := echo.New()
	r := RepoMockNoError{}

	req := httptest.NewRequest(
		http.MethodGet,
		"/iris/iris.gif/?id=XXX-ACCOUNT&uid=1352d6e2-8f86-47ab-95dd-84754480bba9&sid=f313e03a-1163-4cbb-8a65-b5d4f59d13ae&ev=pageload&ed=%7B%22foo%22%3A1%7D&dl=http%3A%2F%2F192.168.35.40%3A8080%2Fpage.html%3Fasd&rl=&ts=1567604284981&de=UTF-8&sr=1920x1080&vp=594x786&cd=24&dt=&bn=Firefox%2068&md=false&ua=Mozilla%2F5.0%20(X11%3B%20Ubuntu%3B%20Linux%20x86_64%3B%20rv%3A68.0)%20Gecko%2F20100101%20Firefox%2F68.0&tz=-120&utm=%7B%22utm_source%22%3A%22foo%22%7D&",
		nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewEventHandler(&r)
	// Act

	// Assert
	if assert.NoError(t, h.Ingest(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(
			t, getInvisiblePixel(), rec.Body.String())
	}
}

func TestIngestReturnsProperReponseIfFails(t *testing.T) {

	// Arrange
	e := echo.New()
	r := RepoMockWithError{}

	req := httptest.NewRequest(
		http.MethodGet,
		"/iris/iris.gif/?id=XXX-ACCOUNT&uid=1352d6e2-8f86-47ab-95dd-84754480bba9&sid=f313e03a-1163-4cbb-8a65-b5d4f59d13ae&ev=pageload&ed=%7B%22foo%22%3A1%7D&dl=http%3A%2F%2F192.168.35.40%3A8080%2Fpage.html%3Fasd&rl=&ts=1567604284981&de=UTF-8&sr=1920x1080&vp=594x786&cd=24&dt=&bn=Firefox%2068&md=false&ua=Mozilla%2F5.0%20(X11%3B%20Ubuntu%3B%20Linux%20x86_64%3B%20rv%3A68.0)%20Gecko%2F20100101%20Firefox%2F68.0&tz=-120&utm=%7B%22utm_source%22%3A%22foo%22%7D&",
		nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewEventHandler(&r)
	// Act

	// Assert
	if assert.NoError(t, h.Ingest(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(
			t, getInvisiblePixel(), rec.Body.String())
	}
}
