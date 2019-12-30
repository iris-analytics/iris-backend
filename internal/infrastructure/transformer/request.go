package transformer

import (
	"time"

	"github.com/iris-analytics/iris-backend/internal/domain/entity"
)

// Request is a the object requests will try to bind to
type Request struct {
	AccountID        string `json:"id" form:"id" query:"id"`
	VisitorID        string `json:"uid" form:"uid" query:"uid"`
	SessionID        string `json:"sid" form:"sid" query:"sid"`
	EventType        string `json:"ev" form:"ev" query:"ev"`
	EventData        string `json:"ed" form:"ed" query:"ed"`
	DocumentLocation string `json:"dl" form:"dl" query:"dl"`
	ReferrerLocation string `json:"rl" form:"rl" query:"rl"`
	DocumentEncoding string `json:"de" form:"de" query:"de"`
	ScreenResolution string `json:"sr" form:"sr" query:"sr"`
	ViewPort         string `json:"vp" form:"vp" query:"vp"`
	ColorDepth       uint16 `json:"cd" form:"cd" query:"cd"`
	DocumentTitle    string `json:"dt" form:"dt" query:"dt"`
	BrowserName      string `json:"bn" form:"bn" query:"bn"`
	IsMobileDevice   bool   `json:"md" form:"md" query:"md"`
	UserAgent        string `json:"ua" form:"ua" query:"ua"`
	TimeZoneOffset   int16  `json:"tz" form:"tz" query:"tz"`
	Utm              string `json:"utm" form:"utm" query:"utm"`

	Timestamp time.Time
	IPAddress string
}

// MakeEvent makes an event out of the known data
func (r *Request) MakeEvent() *entity.Event {
	e := entity.NewEvent(
		r.AccountID,
		r.VisitorID,
		r.SessionID,
		r.EventType,
		r.EventData,

		r.DocumentLocation,
		r.ReferrerLocation,
		r.Timestamp,
		r.DocumentEncoding,
		r.ScreenResolution,

		r.ViewPort,
		r.ColorDepth,
		r.DocumentTitle,
		r.BrowserName,
		r.IsMobileDevice,

		r.UserAgent,
		r.TimeZoneOffset,
		r.Utm,
		r.IPAddress,
	)

	return e
}
