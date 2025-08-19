package request

// GetMACD represents request parameters for Moving Average Convergence Divergence technical indicator.
type GetMACD struct {
	APIKey
	Symbol       string `schema:"symbol"`
	Interval     string `schema:"interval"`
	SeriesType   string `schema:"series_type,omitempty"`
	FastPeriod   int    `schema:"fast_period,omitempty"`
	SlowPeriod   int    `schema:"slow_period,omitempty"`
	SignalPeriod int    `schema:"signal_period,omitempty"`
	OutputSize   int    `schema:"outputsize,omitempty"`
	Format       string `schema:"format,omitempty"`
}
