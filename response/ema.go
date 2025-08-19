package response

// EMA represents the response structure for the Exponential Moving Average technical indicator endpoint.
type EMA struct {
	Meta   EMAMeta   `json:"meta"`
	Values []EMAData `json:"values"`
	Status string    `json:"status"`
}

// EMAMeta contains metadata information about the Exponential Moving Average calculation.
type EMAMeta struct {
	Symbol           string       `json:"symbol"`
	Interval         string       `json:"interval"`
	Currency         string       `json:"currency"`
	ExchangeTimezone string       `json:"exchange_timezone"`
	Exchange         string       `json:"exchange"`
	MicCode          string       `json:"mic_code"`
	Type             string       `json:"type"`
	Indicator        EMAIndicator `json:"indicator"`
}

// EMAIndicator contains metadata about the Exponential Moving Average indicator configuration.
type EMAIndicator struct {
	Name       string `json:"name"`
	SeriesType string `json:"series_type"`
	TimePeriod int    `json:"time_period"`
}

// EMAData represents individual Exponential Moving Average data points.
type EMAData struct {
	Datetime string `json:"datetime"`
	EMA      string `json:"ema"`
}
