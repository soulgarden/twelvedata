package request

// GetRecommendations represents request parameters for analyst recommendations.
type GetRecommendations struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	Country  string `schema:"country,omitempty"`
}
