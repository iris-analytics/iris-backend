package clickhouse

import (
	"database/sql"
	"strings"

	"github.com/iris-analytics/iris-backend/internal/domain/entity"
	"github.com/iris-analytics/iris-backend/internal/domain/repository"
)

// EventRepository persists PageViews in ClickHouse
type EventRepository struct {
	connection  *sql.DB
	targetTable string
}

// NewEventRepository creates a new repository
func NewEventRepository(c *sql.DB, targetTable string) repository.EventRepositoryInterface {
	r := &EventRepository{
		connection:  c,
		targetTable: targetTable,
	}
	return r
}

// Persist persist PageViews into ClickHouse
func (r *EventRepository) Persist(e *entity.Event) (*entity.Event, error) {

	insertSQL := `
	INSERT INTO {{target.table}}
			(
				account,
				timestamp,
				event_type,
				visitor_id,
				session_id,

				event_data,
				document_location,
				referrer_location,
				document_encoding,

				screen_resolution,
				view_port,
				color_depth,
				document_title,
				browser_name,

				is_mobile_device,
				user_agent,
				timezone_offset,
				utm,
				ip_address
			)
		VALUES(
			?,?,?,?,? ,?,?,?,? ,?,?,?,?,? ,?,?,?,?,?
		)
	`
	insertSQL = strings.Replace(insertSQL, "{{target.table}}", r.targetTable, 1)

	tx, _ := r.connection.Begin()
	stmt, _ := tx.Prepare(insertSQL)

	if _, err := stmt.Exec(
		e.GetAccountID(),
		e.GetTimestamp(),
		e.GetEventType(),
		e.GetVisitorID(),
		e.GetSessionID(),

		e.GetEventData(),
		e.GetDocumentLocation(),
		e.GetReferrerLocation(),
		e.GetDocumentEncoding(),

		e.GetScreenResolution(),
		e.GetViewPort(),
		e.GetColorDepth(),
		e.GetDocumentTitle(),
		e.GetBrowserName(),

		e.GetIsMobileDevice(),
		e.GetUserAgent(),
		e.GetTimeZoneOffset(),
		e.GetUtm(),
		e.GetIPAddress(),
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return e, nil
}
