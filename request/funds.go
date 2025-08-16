package request

type GetFunds struct {
	ApiKey
	Symbol   string `schema:"symbol,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	Country  string `schema:"country,omitempty"`
	Format   string `schema:"format,omitempty"`
	Page     int    `schema:"page,omitempty"`
}
