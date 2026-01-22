package request

// GetExchangeRate represents request parameters for exchange rate data.
type GetExchangeRate struct {
	APIKey
	Symbol        string `schema:"symbol,omitempty"`
	Date          string `schema:"date,omitempty"`
	Format        string `schema:"format,omitempty"`
	Delimiter     string `schema:"delimiter,omitempty"`
	DecimalPlaces int    `schema:"dp,omitempty"`
	TimeZone      string `schema:"timezone,omitempty"`
}
