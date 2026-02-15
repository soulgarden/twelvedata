package request

// GetETFFullData represents request parameters for ETF full data.
type GetETFFullData struct {
	APIKey
	Symbol        string `schema:"symbol,omitempty"`
	FIGI          string `schema:"figi,omitempty"`
	ISIN          string `schema:"isin,omitempty"`
	CUSIP         string `schema:"cusip,omitempty"`
	Country       string `schema:"country,omitempty"`
	DecimalPlaces int    `schema:"dp,omitempty"`
}
