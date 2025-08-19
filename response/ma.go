package response

// MA represents the Moving Average technical indicator response data.
type MA struct {
	Meta   MAMeta    `json:"meta"`
	Values []MAValue `json:"values"`
	Status string    `json:"status"`
}

// MAMeta represents the metadata for the Moving Average technical indicator response.
type MAMeta struct {
	Symbol           string      `json:"symbol"`
	Interval         string      `json:"interval"`
	Currency         string      `json:"currency"`
	ExchangeTimezone string      `json:"exchange_timezone"`
	Exchange         string      `json:"exchange"`
	MicCode          string      `json:"mic_code"`
	Type             string      `json:"type"`
	Indicator        MAIndicator `json:"indicator"`
}

// MAIndicator contains metadata about the Moving Average indicator configuration.
type MAIndicator struct {
	Name       string `json:"name"`
	SeriesType string `json:"series_type"`
	TimePeriod int    `json:"time_period"`
	MAType     string `json:"ma_type"`
}

// MAValue represents a single data point in the Moving Average technical indicator response.
type MAValue struct {
	Datetime string `json:"datetime"`
	MA       string `json:"ma"`
}
