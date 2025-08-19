package response

// DEMA represents the Double Exponential Moving Average technical indicator response data.
type DEMA struct {
	Meta   DEMAMeta    `json:"meta"`
	Values []DEMAValue `json:"values"`
	Status string      `json:"status"`
}

// DEMAMeta represents the metadata for the Double Exponential Moving Average technical indicator response.
type DEMAMeta struct {
	Symbol           string        `json:"symbol"`
	Interval         string        `json:"interval"`
	Currency         string        `json:"currency"`
	ExchangeTimezone string        `json:"exchange_timezone"`
	Exchange         string        `json:"exchange"`
	MicCode          string        `json:"mic_code"`
	Type             string        `json:"type"`
	Indicator        DEMAIndicator `json:"indicator"`
}

// DEMAIndicator contains metadata about the Double Exponential Moving Average indicator configuration.
type DEMAIndicator struct {
	Name       string `json:"name"`
	SeriesType string `json:"series_type"`
	TimePeriod int    `json:"time_period"`
}

// DEMAValue represents a single data point in the Double Exponential Moving Average technical indicator response.
type DEMAValue struct {
	Datetime string `json:"datetime"`
	DEMA     string `json:"dema"`
}
