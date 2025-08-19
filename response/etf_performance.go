package response

import "github.com/guregu/null/v6"

// ETFPerformance represents performance metrics for an ETF.
type ETFPerformance struct {
	TrailingReturns TrailingReturns `json:"trailing_returns"`
	AnnualReturns   []AnnualReturn  `json:"annual_returns"`
}

// TrailingReturns represents trailing return percentages over various periods.
type TrailingReturns struct {
	OneDay      null.Float `json:"1d"`
	FiveDays    null.Float `json:"5d"`
	OneMonth    null.Float `json:"1m"`
	ThreeMonths null.Float `json:"3m"`
	SixMonths   null.Float `json:"6m"`
	OneYear     null.Float `json:"1y"`
	ThreeYears  null.Float `json:"3y"`
	FiveYears   null.Float `json:"5y"`
	TenYears    null.Float `json:"10y"`
}

// AnnualReturn represents annual performance data for a specific year.
type AnnualReturn struct {
	Year   int        `json:"year"`
	Return null.Float `json:"return"`
}
