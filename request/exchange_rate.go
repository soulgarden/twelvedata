package request

type GetExchangeRate struct {
	ApiKey
	Symbol        string `schema:"symbol,omitempty"`
	Date          string `schema:"date,omitempty"`
	TimeZone      string `schema:"time_zone,omitempty"`
	DecimalPlaces int    `schema:"dp,omitempty"`
}
