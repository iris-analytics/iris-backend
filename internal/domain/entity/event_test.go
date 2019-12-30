package entity

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EventTestSuite struct {
	suite.Suite
}

func (suite *EventTestSuite) SetupTest() {
	fmt.Sprintln("Before each Event test")
}

func (suite *EventTestSuite) TestGettersReturnProperties() {
	// Arrange
	now := time.Now()

	e := NewEvent(
		"accountID",
		"visitorID",
		"sessionID",
		"eventType",
		"eventData",
		"documentLocation",
		"referrerLocation",
		now,
		"documentEncoding",
		"screenResolution",
		"viewPort",
		24,
		"documentTitle",
		"browserName",
		true,
		"userAgent",
		-120,
		"utm",
		"127.0.0.1",
	)
	// Act

	// Assert
	assert.Equal(suite.T(), "accountID", e.GetAccountID())
	assert.Equal(suite.T(), "visitorID", e.GetVisitorID())
	assert.Equal(suite.T(), "sessionID", e.GetSessionID())
	assert.Equal(suite.T(), "eventType", e.GetEventType())
	assert.Equal(suite.T(), "eventData", e.GetEventData())
	assert.Equal(suite.T(), "documentLocation", e.GetDocumentLocation())
	assert.Equal(suite.T(), "referrerLocation", e.GetReferrerLocation())
	assert.Equal(suite.T(), now, e.GetTimestamp())
	assert.Equal(suite.T(), "documentEncoding", e.GetDocumentEncoding())
	assert.Equal(suite.T(), "screenResolution", e.GetScreenResolution())
	assert.Equal(suite.T(), "viewPort", e.GetViewPort())
	assert.Equal(suite.T(), uint16(24), e.GetColorDepth())
	assert.Equal(suite.T(), "documentTitle", e.GetDocumentTitle())
	assert.Equal(suite.T(), "browserName", e.GetBrowserName())
	assert.Equal(suite.T(), true, e.GetIsMobileDevice())
	assert.Equal(suite.T(), "userAgent", e.GetUserAgent())
	assert.Equal(suite.T(), int16(-120), e.GetTimeZoneOffset())
	assert.Equal(suite.T(), "utm", e.GetUtm())
	assert.Equal(suite.T(), "127.0.0.1", e.GetIPAddress())
}

func TestEventTestSuite(t *testing.T) {
	suite.Run(t, new(EventTestSuite))
}
