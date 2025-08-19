package response

import "github.com/guregu/null/v6"

// Earnings represents the response structure for earnings data grouped by date.
type Earnings struct {
	Earnings map[string][]*Earning `json:"earnings"`
	Status   string                `json:"status"`
}

// Earning represents a single earnings report with estimates and actuals.
type Earning struct {
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
