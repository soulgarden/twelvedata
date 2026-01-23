package request

// GetRecommendations represents request parameters for analyst recommendations.
type GetRecommendations struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Figi     string `schema:"figi,omitempty"`
	Isin     string `schema:"isin,omitempty"`
	Cusip    string `schema:"cusip,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	Country  string `schema:"country,omitempty"`
}
