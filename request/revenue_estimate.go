package request

// GetRevenueEstimate represents request parameters for revenue estimates.
type GetRevenueEstimate struct {
	APIKey
	Symbol        string `schema:"symbol,omitempty"`
	Figi          string `schema:"figi,omitempty"`
	Isin          string `schema:"isin,omitempty"`
	Cusip         string `schema:"cusip,omitempty"`
	Exchange      string `schema:"exchange,omitempty"`
	Country       string `schema:"country,omitempty"`
	DecimalPlaces int    `schema:"dp,omitempty"`
}
