package response

import "github.com/guregu/null/v6"

// EarningsCalendar represents the response structure for earnings calendar data grouped by date.
type EarningsCalendar struct {
	Earnings map[string][]*EarningsCalendarItem `json:"earnings"`
	Status   string                             `json:"status"`
}

// EarningsCalendarItem represents a single earnings report with estimates and actuals.
type EarningsCalendarItem struct {
	Symbol      string     `json:"symbol"`
	Name        string     `json:"name"`
	Currency    string     `json:"currency"`
	Exchange    string     `json:"exchange"`
	MicCode     string     `json:"mic_code"`
	Country     string     `json:"country"`
	Time        string     `json:"time"`
	EPSEstimate null.Float `json:"eps_estimate"`
	EPSActual   null.Float `json:"eps_actual"`
	Difference  null.Float `json:"difference"`
	SurprisePrc null.Float `json:"surprise_prc"`
}
