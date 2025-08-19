package request

// GetETFFullData represents request parameters for ETF full data.
type GetETFFullData struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Country  string `schema:"country,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
}
