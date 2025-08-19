package request

// GetMutualFundFullData represents request parameters for mutual fund full data.
type GetMutualFundFullData struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Country  string `schema:"country,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
}
