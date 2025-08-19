package response

// SMA represents the response structure for the Simple Moving Average technical indicator endpoint.
type SMA struct {
	Meta   SMAMeta   `json:"meta"`
	Values []SMAData `json:"values"`
	Status string    `json:"status"`
}

// SMAMeta contains metadata information about the Simple Moving Average calculation.
type SMAMeta struct {
	Symbol           string       `json:"symbol"`
	Interval         string       `json:"interval"`
	Currency         string       `json:"currency"`
	ExchangeTimezone string       `json:"exchange_timezone"`
	Exchange         string       `json:"exchange"`
	MicCode          string       `json:"mic_code"`
	Type             string       `json:"type"`
	Indicator        SMAIndicator `json:"indicator"`
}

// SMAIndicator contains metadata about the Simple Moving Average indicator configuration.
type SMAIndicator struct {
	Name       string `json:"name"`
	SeriesType string `json:"series_type"`
	TimePeriod int    `json:"time_period"`
}

// SMAData represents individual Simple Moving Average data points.
type SMAData struct {
	Datetime string `json:"datetime"`
	SMA      string `json:"sma"`
}
