package request

// GetETFPerformance represents request parameters for ETF performance data.
type GetETFPerformance struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Country  string `schema:"country,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
}
