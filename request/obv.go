package request

// GetOBV represents the request parameters for the On Balance Volume technical indicator endpoint.
type GetOBV struct {
	APIKey
	Symbol     string `schema:"symbol"`
	Interval   string `schema:"interval"`
	OutputSize int    `schema:"outputsize,omitempty"`
	Format     string `schema:"format,omitempty"`
	Delimiter  string `schema:"delimiter,omitempty"`
}
