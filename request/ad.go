package request

// GetAD represents the request parameters for the Accumulation/Distribution technical indicator endpoint.
type GetAD struct {
	APIKey
	Symbol     string `schema:"symbol"`
	Interval   string `schema:"interval"`
	OutputSize int    `schema:"outputsize,omitempty"`
	Format     string `schema:"format,omitempty"`
	Delimiter  string `schema:"delimiter,omitempty"`
}
