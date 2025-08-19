package request

// GetMarketState represents request parameters for market state data.
type GetMarketState struct {
	APIKey
	Exchange string `schema:"exchange,omitempty"`
	Code     string `schema:"code,omitempty"`
	Country  string `schema:"country,omitempty"`
}
