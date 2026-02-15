package request

// GetMutualFundPerformance represents request parameters for mutual fund performance data.
type GetMutualFundPerformance struct {
	APIKey
	Symbol        string `schema:"symbol,omitempty"`
	FIGI          string `schema:"figi,omitempty"`
	ISIN          string `schema:"isin,omitempty"`
	CUSIP         string `schema:"cusip,omitempty"`
	Country       string `schema:"country,omitempty"`
	DecimalPlaces int    `schema:"dp,omitempty"`
}
