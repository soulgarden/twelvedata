package request

// GetFunds represents request parameters for funds data.
type GetFunds struct {
	APIKey
	Symbol     string `schema:"symbol,omitempty"`
	Figi       string `schema:"figi,omitempty"`
	Isin       string `schema:"isin,omitempty"`
	Cusip      string `schema:"cusip,omitempty"`
	Cik        string `schema:"cik,omitempty"`
	Exchange   string `schema:"exchange,omitempty"`
	Country    string `schema:"country,omitempty"`
	Format     string `schema:"format,omitempty"`
	Delimiter  string `schema:"delimiter,omitempty"`
	ShowPlan   bool   `schema:"show_plan,omitempty"`
	Page       int    `schema:"page,omitempty"`
	OutputSize int    `schema:"outputsize,omitempty"`
}
