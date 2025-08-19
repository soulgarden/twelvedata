package response

// WMA represents the Weighted Moving Average technical indicator response data.
type WMA struct {
	Meta   WMAMeta    `json:"meta"`
	Values []WMAValue `json:"values"`
	Status string     `json:"status"`
}

// WMAMeta represents the metadata for the Weighted Moving Average technical indicator response.
type WMAMeta struct {
	Symbol           string       `json:"symbol"`
	Interval         string       `json:"interval"`
	Currency         string       `json:"currency"`
	ExchangeTimezone string       `json:"exchange_timezone"`
	Exchange         string       `json:"exchange"`
	MicCode          string       `json:"mic_code"`
	Type             string       `json:"type"`
	Indicator        WMAIndicator `json:"indicator"`
}

// WMAIndicator contains metadata about the Weighted Moving Average indicator configuration.
type WMAIndicator struct {
	Name       string `json:"name"`
	SeriesType string `json:"series_type"`
	TimePeriod int    `json:"time_period"`
}

// WMAValue represents a single data point in the Weighted Moving Average technical indicator response.
type WMAValue struct {
	Datetime string `json:"datetime"`
	WMA      string `json:"wma"`
}
