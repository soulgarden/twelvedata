package response

import "github.com/guregu/null/v6"

// ETFPerformance represents the response structure for ETF performance data.
type ETFPerformance struct {
	ETF    ETFPerformanceData `json:"etf"`
	Status string             `json:"status"`
}

// ETFPerformanceData contains the performance information for an ETF.
type ETFPerformanceData struct {
	Performance ETFWorldPerformance `json:"performance"`
}

// ETFWorldPerformance represents detailed performance of an ETF.
type ETFWorldPerformance struct {
	TrailingReturns    []ETFTrailingReturn    `json:"trailing_returns"`
	AnnualTotalReturns []ETFAnnualTotalReturn `json:"annual_total_returns"`
}

// ETFTrailingReturn represents trailing return data for a specific period.
type ETFTrailingReturn struct {
	Period           string     `json:"period"`
	ShareClassReturn null.Float `json:"share_class_return"`
	CategoryReturn   null.Float `json:"category_return"`
}

// ETFAnnualTotalReturn represents annual total return data for a specific year.
type ETFAnnualTotalReturn struct {
	Year             null.Int   `json:"year"`
	ShareClassReturn null.Float `json:"share_class_return"`
	CategoryReturn   null.Float `json:"category_return"`
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
	Year   null.Int   `json:"year"`
	Return null.Float `json:"return"`
}
