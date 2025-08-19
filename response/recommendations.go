package response

import "github.com/guregu/null/v6"

// Recommendations represents analyst recommendations data.
type Recommendations struct {
	Symbol      string                  `json:"symbol"`
	Exchange    string                  `json:"exchange"`
	Summary     RecommendationSummary   `json:"summary"`
	Breakdown   RecommendationBreakdown `json:"breakdown"`
	History     []RecommendationHistory `json:"history"`
	LastUpdated string                  `json:"last_updated"`
}

// RecommendationSummary represents overall recommendation summary.
type RecommendationSummary struct {
	Score          null.Float `json:"score"`
	Recommendation string     `json:"recommendation"`
	AnalystCount   int        `json:"analyst_count"`
}

// RecommendationBreakdown represents detailed breakdown of recommendations.
type RecommendationBreakdown struct {
	StrongBuy  int `json:"strong_buy"`
	Buy        int `json:"buy"`
	Hold       int `json:"hold"`
	Sell       int `json:"sell"`
	StrongSell int `json:"strong_sell"`
}

// RecommendationHistory represents historical recommendation data.
type RecommendationHistory struct {
	Date           string `json:"date"`
	Recommendation string `json:"recommendation"`
	AnalystCount   int    `json:"analyst_count"`
}
