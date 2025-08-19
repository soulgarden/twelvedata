package request

// GetExchanges represents request parameters for exchanges data.
type GetExchanges struct {
	APIKey
	InstrumentType string `schema:"type,omitempty"`
	Name           string `schema:"name,omitempty"`
	Code           string `schema:"code,omitempty"`
	Country        string `schema:"country,omitempty"`
	ShowPlan       bool   `schema:"show_plan,omitempty"`
}
