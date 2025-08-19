package response

// PercentB represents the response structure for %B (Percent B) technical indicator.
type PercentB struct {
	Meta   PercentBMeta   `json:"meta"`
	Values []PercentBData `json:"values"`
	Status string         `json:"status"`
}

// PercentBMeta contains metadata for PercentB response.
type PercentBMeta struct {
	Symbol           string            `json:"symbol"`
	Interval         string            `json:"interval"`
	Currency         string            `json:"currency"`
	ExchangeTimezone string            `json:"exchange_timezone"`
	Exchange         string            `json:"exchange"`
	MicCode          string            `json:"mic_code"`
	Type             string            `json:"type"`
	Indicator        PercentBIndicator `json:"indicator"`
}

// PercentBIndicator contains %B indicator configuration.
type PercentBIndicator struct {
	Name       string `json:"name"`
	SeriesType string `json:"series_type"`
	TimePeriod int    `json:"time_period"`
	SD         int    `json:"sd"`
	MAType     string `json:"ma_type"`
}

// PercentBData represents individual %B data points.
type PercentBData struct {
	Datetime string `json:"datetime"`
	PercentB string `json:"percent_b"`
}
