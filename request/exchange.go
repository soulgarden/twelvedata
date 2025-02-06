package request

type GetExchanges struct {
	ApiKey
	InstrumentType string `schema:"type,omitempty"`
	Name           string `schema:"name,omitempty"`
	Code           string `schema:"code,omitempty"`
	Country        string `schema:"country,omitempty"`
	ShowPlan       bool   `schema:"show_plan,omitempty"`
}
