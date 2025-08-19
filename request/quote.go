package request

// GetQuote represents request parameters for quote data.
type GetQuote struct {
	APIKey
	Symbol           string `schema:"symbol,omitempty"`
	Interval         string `schema:"interval,omitempty"`
	Exchange         string `schema:"exchange,omitempty"`
	MicCode          string `schema:"mic_code,omitempty"`
	Country          string `schema:"country,omitempty"`
	VolumeTimePeriod string `schema:"volume_time_period,omitempty"`
	InstrumentType   string `schema:"type,omitempty"`
	Prepost          string `schema:"prepost,omitempty"`
	Eod              bool   `schema:"eod,omitempty"`
	RollingPeriod    int    `schema:"rolling_period,omitempty"`
	DecimalPlaces    int    `schema:"dp,omitempty"`
	TimeZone         string `schema:"timezone,omitempty"`
}
