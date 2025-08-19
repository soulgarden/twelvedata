package request

// GetMutualFunds represents request parameters for mutual funds directory data.
type GetMutualFunds struct {
	APIKey
	Symbol          string `schema:"symbol,omitempty"`
	Exchange        string `schema:"exchange,omitempty"`
	Country         string `schema:"country,omitempty"`
	Format          string `schema:"format,omitempty"`
	Delimiter       string `schema:"delimiter,omitempty"`
	ShowPlan        bool   `schema:"show_plan,omitempty"`
	Page            int    `schema:"page,omitempty"`
	OutputSize      int    `schema:"outputsize,omitempty"`
	IncludeDelisted bool   `schema:"include_delisted,omitempty"`
}
