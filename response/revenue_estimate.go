package response

import "github.com/guregu/null/v6"

// RevenueEstimate represents revenue estimate data.
type RevenueEstimate struct {
	Symbol      string                     `json:"symbol"`
	Exchange    string                     `json:"exchange"`
	Quarterly   []QuarterlyRevenueEstimate `json:"quarterly"`
	Annual      []AnnualRevenueEstimate    `json:"annual"`
	Trends      RevenueEstimateTrends      `json:"trends"`
	LastUpdated string                     `json:"last_updated"`
}

// QuarterlyRevenueEstimate represents quarterly revenue estimates.
type QuarterlyRevenueEstimate struct {
	Period       string     `json:"period"`
	Year         int        `json:"year"`
	Quarter      int        `json:"quarter"`
	RevEstimate  null.Float `json:"revenue_estimate"`
	RevActual    null.Float `json:"revenue_actual"`
	RevGrowth    null.Float `json:"revenue_growth"`
	AnalystCount int        `json:"analyst_count"`
}

// AnnualRevenueEstimate represents annual revenue estimates.
type AnnualRevenueEstimate struct {
	Year         int        `json:"year"`
	RevEstimate  null.Float `json:"revenue_estimate"`
	RevActual    null.Float `json:"revenue_actual"`
	RevGrowth    null.Float `json:"revenue_growth"`
	AnalystCount int        `json:"analyst_count"`
}

// RevenueEstimateTrends represents revenue estimate trend analysis.
type RevenueEstimateTrends struct {
	CurrentQuarter RevenueEstimateTrend `json:"current_quarter"`
	NextQuarter    RevenueEstimateTrend `json:"next_quarter"`
	CurrentYear    RevenueEstimateTrend `json:"current_year"`
	NextYear       RevenueEstimateTrend `json:"next_year"`
}

// RevenueEstimateTrend represents individual revenue estimate trend data.
type RevenueEstimateTrend struct {
	RevenueEstimate    null.Float `json:"revenue_estimate"`
	RevenueRevision7d  null.Float `json:"revenue_revision_7d"`
	RevenueRevision30d null.Float `json:"revenue_revision_30d"`
	RevenueRevision60d null.Float `json:"revenue_revision_60d"`
	RevenueRevision90d null.Float `json:"revenue_revision_90d"`
	RevenueGrowth      null.Float `json:"revenue_growth"`
	AnalystCount       int        `json:"analyst_count"`
}
