package request

type GetIndices struct {
	ApiKey
	Symbol          string `schema:"symbol,omitempty"`
	Country         string `schema:"country,omitempty"`
	ShowPlan        bool   `schema:"show_plan,omitempty"`
	IncludeDelisted bool   `schema:"include_delisted,omitempty"`
}
