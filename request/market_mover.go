package request

// GetMarketMovers represents request parameters for market movers data.
type GetMarketMovers struct {
	APIKey
	Market           string `schema:"-"` // Market goes in URL path, not query params
	Direction        string `schema:"direction,omitempty"`
	OutputSize       int    `schema:"outputsize,omitempty"`
	Country          string `schema:"country,omitempty"`
	PriceGreaterThan string `schema:"price_greater_than,omitempty"`
	DecimalPlaces    string `schema:"dp,omitempty"`
}
