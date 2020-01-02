package usecase

import (
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/iris-analytics/iris-backend/internal/domain/entity"
	"github.com/iris-analytics/iris-backend/internal/infrastructure/provider"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type DummyEventRepository struct {
	recordedEvent entity.Event
}

func (repo *DummyEventRepository) Record(e *entity.Event) (*entity.Event, error) {
	repo.recordedEvent = *e

	return e, nil
}

func (repo *DummyEventRepository) getRecord() entity.Event {
	return repo.recordedEvent
}

func TestExecuteCallsRepositoryWithProperEvent(t *testing.T) {

	// arrange
	logger := provider.GetSugaredLogger("appName")
	dummyEventRepo := DummyEventRepository{}
	useCase := RecordEvent{
		EventRepository: &dummyEventRepo,
		Logger:          logger,
	}
	req := httptest.NewRequest(
		http.MethodGet,
		"/any/?id=XXX-ACCOUNT&uid=1352d6e2-8f86-47ab-95dd-84754480bba9&sid=f313e03a-1163-4cbb-8a65-b5d4f59d13ae&ev=pageload&ed=%7B%22foo%22%3A1%7D&dl=http%3A%2F%2F192.168.35.40%3A8080%2Fpage.html%3Fasd&rl=&ts=1567604284981&de=UTF-8&sr=1920x1080&vp=594x786&cd=24&dt=&bn=Firefox%2068&md=false&ua=Mozilla%2F5.0%20(X11%3B%20Ubuntu%3B%20Linux%20x86_64%3B%20rv%3A68.0)%20Gecko%2F20100101%20Firefox%2F68.0&tz=-120&utm=%7B%22utm_source%22%3A%22foo%22%7D&",
		strings.NewReader("fooo"),
	)
	rec := httptest.NewRecorder()
	e := echo.New()
	context := e.NewContext(req, rec)
	// act
	err := useCase.Execute(context)
	// assert
	assert.NoError(t, err)
	//time.Sleep(2 * time.Second)

	record := dummyEventRepo.getRecord()

	assert.Equal(t, "XXX-ACCOUNT", record.GetAccountID())
	assert.Equal(t, "1352d6e2-8f86-47ab-95dd-84754480bba9", record.GetVisitorID())
	assert.Equal(t, "f313e03a-1163-4cbb-8a65-b5d4f59d13ae", record.GetSessionID())
	assert.Equal(t, "pageload", record.GetEventType())

	assert.Equal(t, "{\"foo\":1}", record.GetEventData())
	assert.Equal(t, "http://192.168.35.40:8080/page.html?asd", record.GetDocumentLocation())
	assert.Equal(t, "", record.GetReferrerLocation())
	assert.Equal(t, "UTF-8", record.GetDocumentEncoding())

	assert.Equal(t, "1920x1080", record.GetScreenResolution())
	assert.Equal(t, "594x786", record.GetViewPort())
	assert.Equal(t, "", record.GetDocumentTitle())
	assert.Equal(t, "Firefox 68", record.GetBrowserName())

	assert.Equal(t, false, record.GetIsMobileDevice())
	assert.Equal(t, "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0", record.GetUserAgent())
	assert.Equal(t, "{\"utm_source\":\"foo\"}", record.GetUtm())
	assert.Equal(t, "192.0.2.1", record.GetIPAddress())
}
