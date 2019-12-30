package entity

import "time"

// Event contains data about a page view
type Event struct {
	accountID        string
	visitorID        string
	sessionID        string
	eventType        string
	eventData        string
	documentLocation string
	referrerLocation string
	timeStamp        time.Time
	documentEncoding string
	screenResolution string
	viewPort         string
	colorDepth       uint16
	documentTitle    string
	browserName      string
	isMobileDevice   bool
	userAgent        string
	timeZoneOffset   int16 // Minutes away from UTC
	utm              string
	ipAddress        string
}

// GetAccountID gets account
func (e *Event) GetAccountID() string {
	return e.accountID
}

// GetVisitorID gets account
func (e *Event) GetVisitorID() string {
	return e.visitorID
}

// GetSessionID gets account
func (e *Event) GetSessionID() string {
	return e.sessionID
}

// GetEventType gets account
func (e *Event) GetEventType() string {
	return e.eventType
}

// GetEventData gets account
func (e *Event) GetEventData() string {
	return e.eventData
}

// GetDocumentLocation gets account
func (e *Event) GetDocumentLocation() string {
	return e.documentLocation
}

// GetReferrerLocation gets account
func (e *Event) GetReferrerLocation() string {
	return e.referrerLocation
}

// GetTimestamp gets account
func (e *Event) GetTimestamp() time.Time {
	return e.timeStamp
}

// GetDocumentEncoding gets account
func (e *Event) GetDocumentEncoding() string {
	return e.documentEncoding
}

// GetScreenResolution gets account
func (e *Event) GetScreenResolution() string {
	return e.screenResolution
}

// GetViewPort gets account
func (e *Event) GetViewPort() string {
	return e.viewPort
}

// GetColorDepth gets account
func (e *Event) GetColorDepth() uint16 {
	return e.colorDepth
}

// GetDocumentTitle gets account
func (e *Event) GetDocumentTitle() string {
	return e.documentTitle
}

// GetBrowserName gets account
func (e *Event) GetBrowserName() string {
	return e.browserName
}

// GetIsMobileDevice gets account
func (e *Event) GetIsMobileDevice() bool {
	return e.isMobileDevice
}

// GetUserAgent gets account
func (e *Event) GetUserAgent() string {
	return e.userAgent
}

// GetTimeZoneOffset gets account
func (e *Event) GetTimeZoneOffset() int16 {
	return e.timeZoneOffset
}

// GetUtm gets account
func (e *Event) GetUtm() string {
	return e.utm
}

// GetIPAddress gets account
func (e *Event) GetIPAddress() string {
	return e.ipAddress
}

// NewEvent creates new event instance
func NewEvent(
	accountID string,
	visitorID string,
	sessionID string,
	eventType string,
	eventData string,
	documentLocation string,
	referrerLocation string,
	timeStamp time.Time,
	documentEncoding string,
	screenResolution string,
	viewPort string,
	colorDepth uint16,
	documentTitle string,
	browserName string,
	isMobileDevice bool,
	userAgent string,
	timeZoneOffset int16, // Minutes away from UTC
	utm string,
	ipAddress string,
) *Event {

	e := Event{
		accountID:        accountID,
		visitorID:        visitorID,
		sessionID:        sessionID,
		eventType:        eventType,
		eventData:        eventData,
		documentLocation: documentLocation,
		referrerLocation: referrerLocation,
		timeStamp:        timeStamp,
		documentEncoding: documentEncoding,
		screenResolution: screenResolution,
		viewPort:         viewPort,
		colorDepth:       colorDepth,
		documentTitle:    documentTitle,
		browserName:      browserName,
		isMobileDevice:   isMobileDevice,
		userAgent:        userAgent,
		timeZoneOffset:   timeZoneOffset,
		utm:              utm,
		ipAddress:        ipAddress,
	}

	return &e
}
