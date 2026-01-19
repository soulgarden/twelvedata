package request

// GetStock represents request parameters for stock data.
type GetStock struct {
	APIKey
	Symbol          string `schema:"symbol,omitempty"`
	Figi            string `schema:"figi,omitempty"`
	Isin            string `schema:"isin,omitempty"`
	Cusip           string `schema:"cusip,omitempty"`
	Cik             string `schema:"cik,omitempty"`
	Exchange        string `schema:"exchange,omitempty"`
	MicCode         string `schema:"mic_code,omitempty"`
	Country         string `schema:"country,omitempty"`
	InstrumentType  string `schema:"type,omitempty"`
	Format          string `schema:"format,omitempty"`
	Delimiter       string `schema:"delimiter,omitempty"`
	ShowPlan        bool   `schema:"show_plan,omitempty"`
	IncludeDelisted bool   `schema:"include_delisted,omitempty"`
}
