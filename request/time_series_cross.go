package request

// GetTimeSeriesCross represents request parameters for cross time series data.
type GetTimeSeriesCross struct {
	APIKey
	Base          string `schema:"base,omitempty"`
	Quote         string `schema:"quote,omitempty"`
	Interval      string `schema:"interval,omitempty"`
	BaseType      string `schema:"base_type,omitempty"`
	BaseExchange  string `schema:"base_exchange,omitempty"`
	BaseMicCode   string `schema:"base_mic_code,omitempty"`
	QuoteType     string `schema:"quote_type,omitempty"`
	QuoteExchange string `schema:"quote_exchange,omitempty"`
	QuoteMicCode  string `schema:"quote_mic_code,omitempty"`
	OutputSize    int    `schema:"outputsize,omitempty"`
	Format        string `schema:"format,omitempty"`
	Delimiter     string `schema:"delimiter,omitempty"`
	PrePost       bool   `schema:"prepost,omitempty"`
	StartDate     string `schema:"start_date,omitempty"`
	EndDate       string `schema:"end_date,omitempty"`
	Adjust        bool   `schema:"adjust,omitempty"`
	DecimalPlaces int    `schema:"dp,omitempty"`
	TimeZone      string `schema:"timezone,omitempty"`
}
