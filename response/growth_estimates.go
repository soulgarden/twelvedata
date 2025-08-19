package response

import "github.com/guregu/null/v6"

// GrowthEstimates represents growth estimates data.
type GrowthEstimates struct {
	Symbol      string                  `json:"symbol"`
	Exchange    string                  `json:"exchange"`
	EPS         GrowthEstimatesEPS      `json:"eps"`
	Revenue     GrowthEstimatesRevenue  `json:"revenue"`
	Earnings    GrowthEstimatesEarnings `json:"earnings"`
	PEG         GrowthEstimatesPEG      `json:"peg"`
	LastUpdated string                  `json:"last_updated"`
}

// GrowthEstimatesEPS represents EPS growth estimates.
type GrowthEstimatesEPS struct {
	CurrentQuarter null.Float `json:"current_quarter"`
	NextQuarter    null.Float `json:"next_quarter"`
	CurrentYear    null.Float `json:"current_year"`
	NextYear       null.Float `json:"next_year"`
	Next5Years     null.Float `json:"next_5_years"`
	PastYear       null.Float `json:"past_year"`
	Past3Years     null.Float `json:"past_3_years"`
	Past5Years     null.Float `json:"past_5_years"`
	AnalystCount   int        `json:"analyst_count"`
}

// GrowthEstimatesRevenue represents revenue growth estimates.
type GrowthEstimatesRevenue struct {
	CurrentQuarter null.Float `json:"current_quarter"`
	NextQuarter    null.Float `json:"next_quarter"`
	CurrentYear    null.Float `json:"current_year"`
	NextYear       null.Float `json:"next_year"`
	Next5Years     null.Float `json:"next_5_years"`
	PastYear       null.Float `json:"past_year"`
	Past3Years     null.Float `json:"past_3_years"`
	Past5Years     null.Float `json:"past_5_years"`
	AnalystCount   int        `json:"analyst_count"`
}

// GrowthEstimatesEarnings represents earnings growth estimates.
type GrowthEstimatesEarnings struct {
	CurrentYear      null.Float `json:"current_year"`
	NextYear         null.Float `json:"next_year"`
	Next5Years       null.Float `json:"next_5_years"`
	LongTermGrowth   null.Float `json:"long_term_growth"`
	HistoricalGrowth null.Float `json:"historical_growth"`
	AnalystCount     int        `json:"analyst_count"`
}

// GrowthEstimatesPEG represents PEG ratio estimates based on growth.
type GrowthEstimatesPEG struct {
	PEGRatio        null.Float `json:"peg_ratio"`
	PERatio         null.Float `json:"pe_ratio"`
	GrowthRate      null.Float `json:"growth_rate"`
	PEGRatio5Year   null.Float `json:"peg_ratio_5_year"`
	PEGRatioForward null.Float `json:"peg_ratio_forward"`
	AnalystCount    int        `json:"analyst_count"`
}
