package request

// GetMutualFundPerformance represents request parameters for mutual fund performance data.
type GetMutualFundPerformance struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Country  string `schema:"country,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
}
