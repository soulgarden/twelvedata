package request

// GetMutualFundComposition represents request parameters for mutual fund composition data.
type GetMutualFundComposition struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Country  string `schema:"country,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
}
