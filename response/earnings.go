package response

import "github.com/guregu/null/v6"

// Earnings represents the response structure for detailed earnings data.
type Earnings struct {
	Meta     EarningsMeta   `json:"meta"`
	Earnings []EarningsItem `json:"earnings"`
	Status   string         `json:"status"`
}

// EarningsMeta contains metadata for earnings data.
type EarningsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

// EarningsItem represents a single earnings data point with estimates and actuals.
type EarningsItem struct {
	Date        string     `json:"date"`
	Time        string     `json:"time"`
	EPSEstimate null.Float `json:"eps_estimate"`
	EPSActual   null.Float `json:"eps_actual"`
	Difference  null.Float `json:"difference"`
	SurprisePrc null.Float `json:"surprise_prc"`
}
