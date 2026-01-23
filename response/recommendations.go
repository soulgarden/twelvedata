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
	StrongBuy  int `json:"strong_buy"`
	Buy        int `json:"buy"`
	Hold       int `json:"hold"`
	Sell       int `json:"sell"`
	StrongSell int `json:"strong_sell"`
}
