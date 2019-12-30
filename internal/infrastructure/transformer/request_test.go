package transformer

import (
	"testing"
	"time"

	"github.com/iris-analytics/iris-backend/internal/domain/entity"

	"github.com/stretchr/testify/assert"
)

func TestMakeEventReturnsEventObject(t *testing.T) {
	// Arrange

	// Act
	r := Request{
		AccountID:        "myAccount",
		VisitorID:        "myVisitorID",
		SessionID:        "mySessionID",
		EventType:        "myEventType",
		EventData:        "myEventData",
		DocumentLocation: "myDocumentLocation",
		ReferrerLocation: "myReferrerLocation",
		DocumentEncoding: "myDocumentEncoding",
		ScreenResolution: "myScreenResolution",
		ViewPort:         "myViewPort",
		ColorDepth:       24,
		DocumentTitle:    "myDocumentTitle",
		BrowserName:      "myBrowserName",
		IsMobileDevice:   true,
		UserAgent:        "myUserAgent",
		TimeZoneOffset:   30,
		Utm:              "myUtm",

		Timestamp: time.Now(),
		IPAddress: "11.22.33.44",
	}
	e := r.MakeEvent()

	// Assert
	assert.IsType(t, &entity.Event{}, e)
}
