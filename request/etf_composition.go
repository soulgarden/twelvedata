package request

// GetETFComposition represents request parameters for ETF composition data.
type GetETFComposition struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Country  string `schema:"country,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
}
