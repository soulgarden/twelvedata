package request

// GetEarnings represents request parameters for earnings data.
type GetEarnings struct {
	APIKey
	Symbol        string `schema:"symbol,omitempty"`
	Figi          string `schema:"figi,omitempty"`
	Isin          string `schema:"isin,omitempty"`
	Cusip         string `schema:"cusip,omitempty"`
	Exchange      string `schema:"exchange,omitempty"`
	MicCode       string `schema:"mic_code,omitempty"`
	Country       string `schema:"country,omitempty"`
	Type          string `schema:"type,omitempty"`
	Period        string `schema:"period,omitempty"`
	OutputSize    int    `schema:"outputsize,omitempty"`
	Format        string `schema:"format,omitempty"`
	Delimiter     string `schema:"delimiter,omitempty"`
	DecimalPlaces int    `schema:"dp,omitempty"`
	StartDate     string `schema:"start_date,omitempty"`
	EndDate       string `schema:"end_date,omitempty"`
}
