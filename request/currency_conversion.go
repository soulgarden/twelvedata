package request

// GetCurrencyConversion represents request parameters for currency conversion.
type GetCurrencyConversion struct {
	APIKey
	Symbol        string `schema:"symbol,omitempty"`
	Amount        string `schema:"amount,omitempty"`
	Date          string `schema:"date,omitempty"`
	TimeZone      string `schema:"time_zone,omitempty"`
	DecimalPlaces int    `schema:"dp,omitempty"`
}
