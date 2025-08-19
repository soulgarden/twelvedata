package request

// GetPercentB represents request parameters for %B (Percent B) technical indicator.
type GetPercentB struct {
	APIKey
	Symbol             string `schema:"symbol"`
	Interval           string `schema:"interval"`
	SeriesType         string `schema:"series_type,omitempty"`
	TimePeriod         int    `schema:"time_period,omitempty"`
	StandardDeviations int    `schema:"sd,omitempty"`
	MAType             string `schema:"ma_type,omitempty"`
	OutputSize         int    `schema:"outputsize,omitempty"`
	Format             string `schema:"format,omitempty"`
}
