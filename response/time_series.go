package response

// TimeSeries represents time series data response with metadata and values.
type TimeSeries struct {
	Meta   TimeSeriesMeta    `json:"meta"`
	Values []TimeSeriesValue `json:"values"`
	Status string            `json:"status"`
}

// TimeSeriesMeta contains metadata information for time series data.
type TimeSeriesMeta struct {
	Symbol           string `json:"symbol"`
	Interval         string `json:"interval"`
	Currency         string `json:"currency"`
	ExchangeTimezone string `json:"exchange_timezone"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	Type             string `json:"type"`
}

// TimeSeriesValue represents a single time series data point with OHLCV values.
type TimeSeriesValue struct {
	Datetime string `json:"datetime"`
	Open     string `json:"open"`
	High     string `json:"high"`
	Low      string `json:"low"`
	Close    string `json:"close"`
	Volume   string `json:"volume"`
}
