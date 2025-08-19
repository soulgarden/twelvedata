package request

// GetSMA represents the request parameters for the Simple Moving Average technical indicator endpoint.
type GetSMA struct {
	APIKey
	Symbol     string `schema:"symbol"`
	Interval   string `schema:"interval"`
	SeriesType string `schema:"series_type,omitempty"`
	TimePeriod int    `schema:"time_period,omitempty"`
	OutputSize int    `schema:"outputsize,omitempty"`
	Format     string `schema:"format,omitempty"`
	Delimiter  string `schema:"delimiter,omitempty"`
}
