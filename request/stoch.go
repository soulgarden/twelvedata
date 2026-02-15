package request

// GetStoch represents request parameters for Stochastic Oscillator technical indicator.
type GetStoch struct {
	APIKey
	Symbol        string `schema:"symbol,omitempty"`
	FIGI          string `schema:"figi,omitempty"`
	ISIN          string `schema:"isin,omitempty"`
	CUSIP         string `schema:"cusip,omitempty"`
	Interval      string `schema:"interval,omitempty"`
	Exchange      string `schema:"exchange,omitempty"`
	MICCode       string `schema:"mic_code,omitempty"`
	Country       string `schema:"country,omitempty"`
	FastKPeriod   int    `schema:"fast_k_period,omitempty"`
	SlowKPeriod   int    `schema:"slow_k_period,omitempty"`
	SlowDPeriod   int    `schema:"slow_d_period,omitempty"`
	SlowKMAType   string `schema:"slow_kma_type,omitempty"`
	SlowDMAType   string `schema:"slow_dma_type,omitempty"`
	Type          string `schema:"type,omitempty"`
	OutputSize    int    `schema:"outputsize,omitempty"`
	Format        string `schema:"format,omitempty"`
	Delimiter     string `schema:"delimiter,omitempty"`
	Prepost       bool   `schema:"prepost,omitempty"`
	DP            int    `schema:"dp,omitempty"`
	Order         string `schema:"order,omitempty"`
	IncludeOHLC   bool   `schema:"include_ohlc,omitempty"`
	Timezone      string `schema:"timezone,omitempty"`
	Date          string `schema:"date,omitempty"`
	StartDate     string `schema:"start_date,omitempty"`
	EndDate       string `schema:"end_date,omitempty"`
	PreviousClose bool   `schema:"previous_close,omitempty"`
	Adjust        string `schema:"adjust,omitempty"`
}
