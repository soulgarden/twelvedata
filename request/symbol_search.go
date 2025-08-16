package request

type GetSymbolSearch struct {
	ApiKey
	Symbol     string `schema:"symbol"`
	OutputSize int    `schema:"outputsize,omitempty"`
	ShowPlan   bool   `schema:"show_plan,omitempty"`
}
