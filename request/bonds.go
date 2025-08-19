package request

// GetBonds represents request parameters for bonds data.
type GetBonds struct {
	APIKey
	Symbol     string `schema:"symbol,omitempty"`
	Exchange   string `schema:"exchange,omitempty"`
	Country    string `schema:"country,omitempty"`
	Format     string `schema:"format,omitempty"`
	ShowPlan   bool   `schema:"show_plan,omitempty"`
	Page       int    `schema:"page,omitempty"`
	OutputSize int    `schema:"outputsize,omitempty"`
}
