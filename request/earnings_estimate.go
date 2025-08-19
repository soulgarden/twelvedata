package request

// GetEarningsEstimate represents request parameters for earnings estimates.
type GetEarningsEstimate struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	Country  string `schema:"country,omitempty"`
}
