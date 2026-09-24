package request

// GetETFs represents request parameters for ETFs catalog data.
type GetETFs struct {
	APIKey
	Symbol          string `schema:"symbol,omitempty"`
	FIGI            string `schema:"figi,omitempty"`
	ISIN            string `schema:"isin,omitempty"`
	CUSIP           string `schema:"cusip,omitempty"`
	CIK             string `schema:"cik,omitempty"`
	Exchange        string `schema:"exchange,omitempty"`
	MicCode         string `schema:"mic_code,omitempty"`
	Country         string `schema:"country,omitempty"`
	Format          string `schema:"format,omitempty"`
	Delimiter       string `schema:"delimiter,omitempty"`
	ShowPlan        bool   `schema:"show_plan,omitempty"`
	IncludeDelisted bool   `schema:"include_delisted,omitempty"`
	// OutputSize limits the catalog page. Omit it to request all available records.
	OutputSize int `schema:"outputsize,omitempty"`
	// Page selects a result page; the provider defaults to 1.
	Page int `schema:"page,omitempty"`
}
