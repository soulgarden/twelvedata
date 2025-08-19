package request

// GetPriceTarget represents request parameters for analyst price targets.
type GetPriceTarget struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	Country  string `schema:"country,omitempty"`
}
