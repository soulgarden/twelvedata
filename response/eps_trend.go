package response

import "github.com/guregu/null/v6"

// EPSTrend represents EPS trend analysis data.
type EPSTrend struct {
	Meta     AnalysisMeta    `json:"meta"`
	EPSTrend []EPSTrendEntry `json:"eps_trend"`
	Status   string          `json:"status"`
}

// EPSTrendEntry represents a single EPS trend record.
type EPSTrendEntry struct {
	Date            string     `json:"date"`
	Period          string     `json:"period"`
	CurrentEstimate null.Float `json:"current_estimate"`
	SevenDaysAgo    null.Float `json:"7_days_ago"`
	ThirtyDaysAgo   null.Float `json:"30_days_ago"`
	SixtyDaysAgo    null.Float `json:"60_days_ago"`
	NinetyDaysAgo   null.Float `json:"90_days_ago"`
}
