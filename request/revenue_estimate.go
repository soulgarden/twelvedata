package request

// GetRevenueEstimate represents request parameters for revenue estimates.
type GetRevenueEstimate struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	Country  string `schema:"country,omitempty"`
	Format   string `schema:"format,omitempty"`
}
