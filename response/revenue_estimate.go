package response

import "github.com/guregu/null/v6"

// RevenueEstimate represents revenue estimate data.
type RevenueEstimate struct {
	Meta            AnalysisMeta           `json:"meta"`
	RevenueEstimate []RevenueEstimateEntry `json:"revenue_estimate"`
	Status          string                 `json:"status"`
}

// RevenueEstimateEntry represents a single revenue estimate record.
type RevenueEstimateEntry struct {
	Date             string     `json:"date"`
	Period           string     `json:"period"`
	NumberOfAnalysts int        `json:"number_of_analysts"`
	AvgEstimate      null.Float `json:"avg_estimate"`
	LowEstimate      null.Float `json:"low_estimate"`
	HighEstimate     null.Float `json:"high_estimate"`
	YearAgoSales     null.Float `json:"year_ago_sales"`
	SalesGrowth      null.Float `json:"sales_growth"`
}
