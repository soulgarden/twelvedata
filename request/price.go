package request

type GetPrice struct {
	ApiKey
	Symbol         string `schema:"symbol,omitempty"`
	FIGI           string `schema:"figi,omitempty"`
	ISIN           string `schema:"isin,omitempty"`
	CUSIP          string `schema:"cusip,omitempty"`
	Exchange       string `schema:"exchange,omitempty"`
	MicCode        string `schema:"mic_code,omitempty"`
	Country        string `schema:"country,omitempty"`
	InstrumentType string `schema:"type,omitempty"`
	Format         string `schema:"format,omitempty"`
	Delimiter      string `schema:"delimiter,omitempty"`
	PrePost        bool   `schema:"prepost,omitempty"`
	DecimalPlaces  int    `schema:"dp,omitempty"`
}
