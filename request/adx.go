package request

// GetADX represents the request parameters for the Average Directional Index technical indicator endpoint.
type GetADX struct {
	APIKey
	Symbol     string `schema:"symbol"`
	Interval   string `schema:"interval"`
	TimePeriod int    `schema:"time_period,omitempty"`
	OutputSize int    `schema:"outputsize,omitempty"`
	Format     string `schema:"format,omitempty"`
	Delimiter  string `schema:"delimiter,omitempty"`
}
