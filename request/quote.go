package request

// GetQuote represents request parameters for quote data.
type GetQuote struct {
	APIKey
	Symbol           string `schema:"symbol,omitempty"`
	FIGI             string `schema:"figi,omitempty"`
	ISIN             string `schema:"isin,omitempty"`
	CUSIP            string `schema:"cusip,omitempty"`
	Interval         string `schema:"interval,omitempty"`
	Exchange         string `schema:"exchange,omitempty"`
	MicCode          string `schema:"mic_code,omitempty"`
	Country          string `schema:"country,omitempty"`
	VolumeTimePeriod int    `schema:"volume_time_period,omitempty"`
	InstrumentType   string `schema:"type,omitempty"`
	Format           string `schema:"format,omitempty"`
	Delimiter        string `schema:"delimiter,omitempty"`
	Prepost          bool   `schema:"prepost,omitempty"`
	Eod              bool   `schema:"eod,omitempty"`
	RollingPeriod    int    `schema:"rolling_period,omitempty"`
	DecimalPlaces    int    `schema:"dp,omitempty"`
	TimeZone         string `schema:"timezone,omitempty"`
}
