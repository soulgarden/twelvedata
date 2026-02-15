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

// PathParams returns URL path parameters for the market movers endpoint.
func (req GetMarketMovers) PathParams() map[string]string {
	return map[string]string{
		"market": req.Market,
	}
}
