package response

import "github.com/guregu/null/v6"

// MutualFundPerformance represents the response structure for mutual fund performance data.
type MutualFundPerformance struct {
	MutualFund MutualFundPerformanceData `json:"mutual_fund"`
	Status     string                    `json:"status"`
}

// MutualFundPerformanceData contains the performance information for a mutual fund.
type MutualFundPerformanceData struct {
	Performance MutualFundPerformanceInfo `json:"performance"`
}

// MutualFundPerformanceInfo represents detailed performance of a mutual fund.
type MutualFundPerformanceInfo struct {
	TrailingReturns       []MutualFundTrailingReturn       `json:"trailing_returns"`
	AnnualTotalReturns    []MutualFundAnnualTotalReturn    `json:"annual_total_returns"`
	QuarterlyTotalReturns []MutualFundQuarterlyTotalReturn `json:"quarterly_total_returns"`
	LoadAdjustedReturn    []MutualFundLoadAdjustedReturn   `json:"load_adjusted_return"`
}

// MutualFundTrailingReturn represents trailing return data for a specific period.
type MutualFundTrailingReturn struct {
	Period           string     `json:"period"`
	ShareClassReturn null.Float `json:"share_class_return"`
	CategoryReturn   null.Float `json:"category_return"`
	RankInCategory   null.Int   `json:"rank_in_category"`
}

// MutualFundAnnualTotalReturn represents annual total return data for a specific year.
type MutualFundAnnualTotalReturn struct {
	Year             int        `json:"year"`
	ShareClassReturn null.Float `json:"share_class_return"`
	CategoryReturn   null.Float `json:"category_return"`
}

// MutualFundQuarterlyTotalReturn represents quarterly total return data for a specific year.
type MutualFundQuarterlyTotalReturn struct {
	Year int        `json:"year"`
	Q1   null.Float `json:"q1"`
	Q2   null.Float `json:"q2"`
	Q3   null.Float `json:"q3"`
	Q4   null.Float `json:"q4"`
}

// MutualFundLoadAdjustedReturn represents load adjusted return data for a specific period.
type MutualFundLoadAdjustedReturn struct {
	Period string     `json:"period"`
	Return null.Float `json:"return"`
}
