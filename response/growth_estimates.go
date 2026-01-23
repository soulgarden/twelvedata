package response

import "github.com/guregu/null/v6"

// GrowthEstimates represents growth estimates data.
type GrowthEstimates struct {
	Meta            AnalysisMeta        `json:"meta"`
	GrowthEstimates GrowthEstimatesData `json:"growth_estimates"`
	Status          string              `json:"status"`
}

// GrowthEstimatesData represents growth estimates data points.
type GrowthEstimatesData struct {
	CurrentQuarter null.Float `json:"current_quarter"`
	NextQuarter    null.Float `json:"next_quarter"`
	CurrentYear    null.Float `json:"current_year"`
	NextYear       null.Float `json:"next_year"`
	Next5YearsPA   null.Float `json:"next_5_years_pa"`
	Past5YearsPA   null.Float `json:"past_5_years_pa"`
}
