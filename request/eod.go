package request

type GetEOD struct {
	ApiKey
	Symbol         string `schema:"symbol,omitempty"`
	FIGI           string `schema:"figi,omitempty"`
	ISIN           string `schema:"isin,omitempty"`
	CUSIP          string `schema:"cusip,omitempty"`
	Exchange       string `schema:"exchange,omitempty"`
	MicCode        string `schema:"mic_code,omitempty"`
	Country        string `schema:"country,omitempty"`
	InstrumentType string `schema:"type,omitempty"`
	Date           string `schema:"date,omitempty"`
	PrePost        bool   `schema:"prepost,omitempty"`
	DecimalPlaces  int    `schema:"dp,omitempty"`
}
