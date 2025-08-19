package request

// GetETFTypes represents request parameters for ETF types endpoint.
type GetETFTypes struct {
	APIKey
	Country string `schema:"country,omitempty"`
	Format  string `schema:"format,omitempty"`
}
