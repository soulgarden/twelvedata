package request

// GetBBands represents the request parameters for the Bollinger Bands technical indicator endpoint.
type GetBBands struct {
	APIKey
	Symbol             string `schema:"symbol"`
	Interval           string `schema:"interval"`
	SeriesType         string `schema:"series_type,omitempty"`
	TimePeriod         int    `schema:"time_period,omitempty"`
	StandardDeviations int    `schema:"sd,omitempty"`
	MAType             int    `schema:"ma_type,omitempty"`
	OutputSize         int    `schema:"outputsize,omitempty"`
	Format             string `schema:"format,omitempty"`
	Delimiter          string `schema:"delimiter,omitempty"`
}
