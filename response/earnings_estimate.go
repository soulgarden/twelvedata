package response

import "github.com/guregu/null/v6"

// EarningsEstimate represents earnings estimate data.
type EarningsEstimate struct {
	Symbol      string              `json:"symbol"`
	Exchange    string              `json:"exchange"`
	Quarterly   []QuarterlyEstimate `json:"quarterly"`
	Annual      []AnnualEstimate    `json:"annual"`
	Trends      EarningsTrends      `json:"trends"`
	LastUpdated string              `json:"last_updated"`
}

// QuarterlyEstimate represents quarterly earnings estimates.
type QuarterlyEstimate struct {
	Period       string     `json:"period"`
	Year         int        `json:"year"`
	Quarter      int        `json:"quarter"`
	EPSEstimate  null.Float `json:"eps_estimate"`
	EPSActual    null.Float `json:"eps_actual"`
	RevEstimate  null.Float `json:"revenue_estimate"`
	RevActual    null.Float `json:"revenue_actual"`
	AnalystCount int        `json:"analyst_count"`
}

// AnnualEstimate represents annual earnings estimates.
type AnnualEstimate struct {
	Year         int        `json:"year"`
	EPSEstimate  null.Float `json:"eps_estimate"`
	EPSGrowth    null.Float `json:"eps_growth"`
	RevEstimate  null.Float `json:"revenue_estimate"`
	RevGrowth    null.Float `json:"revenue_growth"`
	AnalystCount int        `json:"analyst_count"`
}

// EarningsTrends represents earnings trend analysis.
type EarningsTrends struct {
	CurrentQuarter EarningsTrend `json:"current_quarter"`
	NextQuarter    EarningsTrend `json:"next_quarter"`
	CurrentYear    EarningsTrend `json:"current_year"`
	NextYear       EarningsTrend `json:"next_year"`
}

// EarningsTrend represents individual earnings trend data.
type EarningsTrend struct {
	EPSEstimate     null.Float `json:"eps_estimate"`
	EPSRevision7d   null.Float `json:"eps_revision_7d"`
	EPSRevision30d  null.Float `json:"eps_revision_30d"`
	EPSRevision60d  null.Float `json:"eps_revision_60d"`
	EPSRevision90d  null.Float `json:"eps_revision_90d"`
	RevenueEstimate null.Float `json:"revenue_estimate"`
	AnalystCount    int        `json:"analyst_count"`
}
