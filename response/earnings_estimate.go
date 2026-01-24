package response

import "github.com/guregu/null/v6"

// EarningsEstimate represents earnings estimate data.
type EarningsEstimate struct {
	Meta             AnalysisMeta            `json:"meta"`
	EarningsEstimate []EarningsEstimateEntry `json:"earnings_estimate"`
	Status           string                  `json:"status"`
}

// EarningsEstimateEntry represents a single earnings estimate record.
type EarningsEstimateEntry struct {
	Date             string     `json:"date"`
	Period           string     `json:"period"`
	NumberOfAnalysts null.Int   `json:"number_of_analysts"`
	AvgEstimate      null.Float `json:"avg_estimate"`
	LowEstimate      null.Float `json:"low_estimate"`
	HighEstimate     null.Float `json:"high_estimate"`
	YearAgoEPS       null.Float `json:"year_ago_eps"`
}
