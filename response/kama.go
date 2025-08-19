package response

// KAMA represents the Kaufman Adaptive Moving Average technical indicator response data.
type KAMA struct {
	Meta   KAMAMeta    `json:"meta"`
	Values []KAMAValue `json:"values"`
	Status string      `json:"status"`
}

// KAMAMeta represents the metadata for the Kaufman Adaptive Moving Average technical indicator response.
type KAMAMeta struct {
	Symbol           string        `json:"symbol"`
	Interval         string        `json:"interval"`
	Currency         string        `json:"currency"`
	ExchangeTimezone string        `json:"exchange_timezone"`
	Exchange         string        `json:"exchange"`
	MicCode          string        `json:"mic_code"`
	Type             string        `json:"type"`
	Indicator        KAMAIndicator `json:"indicator"`
}

// KAMAIndicator contains metadata about the Kaufman Adaptive Moving Average indicator configuration.
type KAMAIndicator struct {
	Name       string `json:"name"`
	SeriesType string `json:"series_type"`
	TimePeriod int    `json:"time_period"`
}

// KAMAValue represents a single data point in the Kaufman Adaptive Moving Average technical indicator response.
type KAMAValue struct {
	Datetime string `json:"datetime"`
	KAMA     string `json:"kama"`
}
