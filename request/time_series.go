package request

// GetTimeSeries represents request parameters for time series data.
type GetTimeSeries struct {
	APIKey
	Symbol         string `schema:"symbol,omitempty"`
	FIGI           string `schema:"figi,omitempty"`
	ISIN           string `schema:"isin,omitempty"`
	CUSIP          string `schema:"cusip,omitempty"`
	Interval       string `schema:"interval,omitempty"`
	Exchange       string `schema:"exchange,omitempty"`
	MicCode        string `schema:"mic_code,omitempty"`
	Country        string `schema:"country,omitempty"`
	InstrumentType string `schema:"type,omitempty"`
	OutputSize     int    `schema:"outputsize,omitempty"`
	Format         string `schema:"format,omitempty"`
	Delimiter      string `schema:"delimiter,omitempty"`
	PrePost        bool   `schema:"prepost,omitempty"`
	DecimalPlaces  int    `schema:"dp,omitempty"`
	Order          string `schema:"order,omitempty"`
	TimeZone       string `schema:"timezone,omitempty"`
	Date           string `schema:"date,omitempty"`
	StartDate      string `schema:"start_date,omitempty"`
	EndDate        string `schema:"end_date,omitempty"`
	PreviousClose  bool   `schema:"previous_close,omitempty"`
	Adjust         string `schema:"adjust,omitempty"`
}
