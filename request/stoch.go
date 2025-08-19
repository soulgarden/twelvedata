package request

// GetStoch represents request parameters for Stochastic Oscillator technical indicator.
type GetStoch struct {
	APIKey
	Symbol      string `schema:"symbol"`
	Interval    string `schema:"interval"`
	FastKPeriod int    `schema:"fast_k_period,omitempty"`
	SlowKPeriod int    `schema:"slow_k_period,omitempty"`
	SlowDPeriod int    `schema:"slow_d_period,omitempty"`
	SlowKMAType string `schema:"slow_kma_type,omitempty"`
	SlowDMAType string `schema:"slow_dma_type,omitempty"`
	OutputSize  int    `schema:"outputsize,omitempty"`
	Format      string `schema:"format,omitempty"`
}
