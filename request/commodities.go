package request

type GetCommodities struct {
	ApiKey
	Symbol    string `schema:"symbol,omitempty"`
	Category  string `schema:"category,omitempty"`
	Format    string `schema:"format,omitempty"`
	Delimiter string `schema:"delimiter,omitempty"`
}
