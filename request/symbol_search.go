package request

// GetSymbolSearch represents request parameters for symbol search data.
type GetSymbolSearch struct {
	APIKey
	Symbol     string `schema:"symbol"`
	OutputSize int    `schema:"outputsize,omitempty"`
	ShowPlan   bool   `schema:"show_plan,omitempty"`
}
