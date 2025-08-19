package request

// GetMutualFundTypes represents request parameters for mutual fund types endpoint.
type GetMutualFundTypes struct {
	APIKey
	Country string `schema:"country,omitempty"`
	Format  string `schema:"format,omitempty"`
}
