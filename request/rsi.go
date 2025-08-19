package request

// GetRSI represents request parameters for Relative Strength Index technical indicator.
type GetRSI struct {
	APIKey
	Symbol     string `schema:"symbol"`
	Interval   string `schema:"interval"`
	SeriesType string `schema:"series_type,omitempty"`
	TimePeriod int    `schema:"time_period,omitempty"`
	OutputSize int    `schema:"outputsize,omitempty"`
	Format     string `schema:"format,omitempty"`
}
