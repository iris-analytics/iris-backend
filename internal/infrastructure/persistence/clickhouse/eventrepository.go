package clickhouse

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"

	"github.com/iris-analytics/iris-backend/internal/domain/entity"
	"github.com/iris-analytics/iris-backend/internal/domain/repository"
	"github.com/sethgrid/pester"
)

// EventRepository persists PageViews in ClickHouse
type EventRepository struct {
	httpClient      pester.Client
	clickHouseDSN   string
	clickHouseTable string
}

// NewEventRepository creates a new repository
func NewEventRepository(httpClient *pester.Client, clickHouseDSN string, clickHouseTable string) repository.EventRepositoryInterface {
	r := &EventRepository{
		httpClient:      *httpClient,
		clickHouseDSN:   clickHouseDSN,
		clickHouseTable: clickHouseTable,
	}
	return r
}

// Persist persist PageViews into ClickHouse
func (r *EventRepository) Persist(e *entity.Event) (*entity.Event, error) {

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

	var buffer bytes.Buffer
	gz := gzip.NewWriter(&buffer)
	_, _ = gz.Write([]byte(CSVInsertLine))
	gz.Close()

	request, _ := http.NewRequest(
		"POST",
		r.clickHouseDSN+"/?query=INSERT%20INTO%20"+r.clickHouseTable+"%20FORMAT%20CSV",
		&buffer,
	)
	request.Header.Set("Content-Encoding", "gzip")

	response, err := r.httpClient.Do(request)

	if err != nil {
		return e, err
	}

	if response.StatusCode != 200 {
		return e, err
	}

	return e, nil
}

func boolToInt8(v bool) int8 {
	if v == true {
		return 1
	}

	return 0
}

func doubleQuoteQuotes(s string) string {
	return strings.Replace(s, "\"", "\"\"", -1)
}
