package response

// TRMA represents the Triangular Moving Average technical indicator response data.
type TRMA struct {
	Meta   TRMAMeta    `json:"meta"`
	Values []TRMAValue `json:"values"`
	Status string      `json:"status"`
}

// TRMAMeta represents the metadata for the Triangular Moving Average technical indicator response.
type TRMAMeta struct {
	Symbol           string        `json:"symbol"`
	Interval         string        `json:"interval"`
	Currency         string        `json:"currency"`
	ExchangeTimezone string        `json:"exchange_timezone"`
	Exchange         string        `json:"exchange"`
	MicCode          string        `json:"mic_code"`
	Type             string        `json:"type"`
	Indicator        TRMAIndicator `json:"indicator"`
}

// TRMAIndicator contains metadata about the Triangular Moving Average indicator configuration.
type TRMAIndicator struct {
	Name       string `json:"name"`
	SeriesType string `json:"series_type"`
	TimePeriod int    `json:"time_period"`
}

// TRMAValue represents a single data point in the Triangular Moving Average technical indicator response.
type TRMAValue struct {
	Datetime string `json:"datetime"`
	TRMA     string `json:"trma"`
}
