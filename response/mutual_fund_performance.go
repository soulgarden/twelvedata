package response

import "github.com/guregu/null/v6"

// MutualFundPerformance represents performance metrics for a mutual fund.
type MutualFundPerformance struct {
	TrailingReturns     TrailingReturns     `json:"trailing_returns"`
	AnnualReturns       []AnnualReturn      `json:"annual_returns"`
	QuarterlyReturns    []QuarterlyReturn   `json:"quarterly_returns"`
	LoadAdjustedReturns LoadAdjustedReturns `json:"load_adjusted_returns"`
}

// QuarterlyReturn represents quarterly performance data.
type QuarterlyReturn struct {
	Quarter string     `json:"quarter"`
	Year    int        `json:"year"`
	Return  null.Float `json:"return"`
}

// LoadAdjustedReturns represents returns adjusted for fund loads.
type LoadAdjustedReturns struct {
	OneYear    null.Float `json:"1y"`
	ThreeYears null.Float `json:"3y"`
	FiveYears  null.Float `json:"5y"`
	TenYears   null.Float `json:"10y"`
}
