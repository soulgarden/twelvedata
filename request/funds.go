package request

// GetFunds represents request parameters for funds data.
type GetFunds struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	Country  string `schema:"country,omitempty"`
	Format   string `schema:"format,omitempty"`
	Page     int    `schema:"page,omitempty"`
}
