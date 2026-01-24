package response

import "github.com/guregu/null/v6"

// Recommendations represents analyst recommendations data.
type Recommendations struct {
	Meta   AnalysisMeta         `json:"meta"`
	Trends RecommendationTrends `json:"trends"`
	Rating null.Float           `json:"rating"`
	Status string               `json:"status"`
}

// RecommendationTrends represents recommendations trend data by month.
type RecommendationTrends struct {
	CurrentMonth   RecommendationTrend `json:"current_month"`
	PreviousMonth  RecommendationTrend `json:"previous_month"`
	TwoMonthsAgo   RecommendationTrend `json:"2_months_ago"`
	ThreeMonthsAgo RecommendationTrend `json:"3_months_ago"`
}

// RecommendationTrend represents recommendation counts for a period.
type RecommendationTrend struct {
	StrongBuy  null.Int `json:"strong_buy"`
	Buy        null.Int `json:"buy"`
	Hold       null.Int `json:"hold"`
	Sell       null.Int `json:"sell"`
	StrongSell null.Int `json:"strong_sell"`
}
