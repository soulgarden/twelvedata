package request

type GetMarketMovers struct {
	ApiKey
	Market           string  `schema:"-"` // Market goes in URL path, not query params
	Direction        string  `schema:"direction,omitempty"`
	OutputSize       int     `schema:"outputsize,omitempty"`
	Country          string  `schema:"country,omitempty"`
	PriceGreaterThan float64 `schema:"price_greater_than,omitempty"`
	DecimalPlaces    int     `schema:"dp,omitempty"`
}
