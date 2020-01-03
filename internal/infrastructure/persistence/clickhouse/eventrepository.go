package clickhouse

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/iris-analytics/iris-backend/internal/domain/entity"
	"github.com/iris-analytics/iris-backend/internal/domain/repository"
	"github.com/sethgrid/pester"
	"go.uber.org/zap"
)

// EventRepository persists events in ClickHouse
type EventRepository struct {
	httpClient      *pester.Client
	clickHouseDSN   string
	clickHouseTable string
	logger          *zap.SugaredLogger
}

// NewEventRepository creates and returns a new instance of EventRepository
func NewEventRepository(
	httpClient *pester.Client,
	clickHouseDSN string,
	clickHouseTable string,
	logger *zap.SugaredLogger,
) repository.EventRepositoryInterface {
	r := &EventRepository{
		httpClient:      httpClient,
		clickHouseDSN:   clickHouseDSN,
		clickHouseTable: clickHouseTable,
		logger:          logger,
	}
	return r
}

// Record event into ClickHouse
// We use CSV since it's fast and it's considerably simpler.
// Being one row at the time it makes no sense to apply gzip.
// The HTTP request to ClickHouse runs in a goroutine therefore the method execution finishes
// before result (or retry) happens, relying on logging in case of error
func (r *EventRepository) Record(e *entity.Event) (*entity.Event, error) {

	CSVInsertLine := fmt.Sprintf(`"%v","%v","%v","%v","%v","%v","%v","%v","%v","%v","%v",%d,"%v","%v",%d,"%v",%d,"%v","%v"`,
		doubleQuoteQuotes(e.GetAccountID()),
		doubleQuoteQuotes(e.GetTimestamp().Format("2006-01-02 15:04:05")),
		doubleQuoteQuotes(e.GetEventType()),
		doubleQuoteQuotes(e.GetVisitorID()),
		doubleQuoteQuotes(e.GetSessionID()),

		doubleQuoteQuotes(e.GetEventData()),
		doubleQuoteQuotes(e.GetDocumentLocation()),
		doubleQuoteQuotes(e.GetReferrerLocation()),
		doubleQuoteQuotes(e.GetDocumentEncoding()),

		doubleQuoteQuotes(e.GetScreenResolution()),
		doubleQuoteQuotes(e.GetViewPort()),
		e.GetColorDepth(),
		doubleQuoteQuotes(e.GetDocumentTitle()),
		doubleQuoteQuotes(e.GetBrowserName()),

		boolToInt8(e.GetIsMobileDevice()),
		doubleQuoteQuotes(e.GetUserAgent()),
		e.GetTimeZoneOffset(),
		doubleQuoteQuotes(e.GetUtm()),
		doubleQuoteQuotes(e.GetIPAddress()),
	)

	request, err := http.NewRequest(
		"POST",
		r.clickHouseDSN+"/?query=INSERT%20INTO%20"+r.clickHouseTable+"%20FORMAT%20CSV",
		strings.NewReader(CSVInsertLine),
	)

	if err != nil {
		return e, err
	}

	go func() {
		_, err := r.httpClient.Do(request)
		if err != nil {
			r.logger.Error(err)
		}
	}()

	return e, nil
}

// ClickHouse doesn't have a bool type so it must be either 0 or 1
func boolToInt8(v bool) int8 {
	if v == true {
		return 1
	}

	return 0
}

// This is required for escaping strings to be inserted as CSV, at least in CH
func doubleQuoteQuotes(s string) string {
	return strings.Replace(s, "\"", "\"\"", -1)
}
