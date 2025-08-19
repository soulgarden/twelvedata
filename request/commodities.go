package request

// GetCommodities represents request parameters for commodities data.
type GetCommodities struct {
	APIKey
	Symbol    string `schema:"symbol,omitempty"`
	Category  string `schema:"category,omitempty"`
	Format    string `schema:"format,omitempty"`
	Delimiter string `schema:"delimiter,omitempty"`
}
