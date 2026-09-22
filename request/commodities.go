package request

// GetCommodities represents request parameters for commodities data.
type GetCommodities struct {
	APIKey
	Symbol    string `schema:"symbol,omitempty"`
	Category  string `schema:"category,omitempty"`
	Format    string `schema:"format,omitempty"`
	Delimiter string `schema:"delimiter,omitempty"`
	// OutputSize limits the catalog page. Omit it to request all available records.
	OutputSize int `schema:"outputsize,omitempty"`
	// Page selects a result page; the provider defaults to 1.
	Page int `schema:"page,omitempty"`
}
