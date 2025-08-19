package request

// GetETFFamilies represents request parameters for ETF families endpoint.
type GetETFFamilies struct {
	APIKey
	Country string `schema:"country,omitempty"`
	Format  string `schema:"format,omitempty"`
}
