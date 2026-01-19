package request

// GetExchanges represents request parameters for exchanges data.
type GetExchanges struct {
	APIKey
	InstrumentType string `schema:"type,omitempty"`
	Name           string `schema:"name,omitempty"`
	Code           string `schema:"code,omitempty"`
	Country        string `schema:"country,omitempty"`
	Format         string `schema:"format,omitempty"`
	Delimiter      string `schema:"delimiter,omitempty"`
	ShowPlan       bool   `schema:"show_plan,omitempty"`
}
