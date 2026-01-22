package request

// GetCurrencyConversion represents request parameters for currency conversion.
type GetCurrencyConversion struct {
	APIKey
	Symbol        string `schema:"symbol,omitempty"`
	Amount        string `schema:"amount,omitempty"`
	Date          string `schema:"date,omitempty"`
	Format        string `schema:"format,omitempty"`
	Delimiter     string `schema:"delimiter,omitempty"`
	DecimalPlaces int    `schema:"dp,omitempty"`
	TimeZone      string `schema:"timezone,omitempty"`
}
