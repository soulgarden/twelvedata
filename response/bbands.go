package response

// BBands represents the response structure for the Bollinger Bands technical indicator endpoint.
type BBands struct {
	Meta   BBandsMeta   `json:"meta"`
	Values []BbandsData `json:"values"`
	Status string       `json:"status"`
}

// BBandsMeta contains metadata information about the Bollinger Bands calculation.
type BBandsMeta struct {
	Symbol           string          `json:"symbol"`
	Interval         string          `json:"interval"`
	Currency         string          `json:"currency"`
	ExchangeTimezone string          `json:"exchange_timezone"`
	Exchange         string          `json:"exchange"`
	MicCode          string          `json:"mic_code"`
	Type             string          `json:"type"`
	Indicator        BbandsIndicator `json:"indicator"`
}

// BbandsIndicator contains metadata about the Bollinger Bands indicator configuration.
type BbandsIndicator struct {
	Name       string `json:"name"`
	SeriesType string `json:"series_type"`
	TimePeriod int    `json:"time_period"`
	SD         int    `json:"sd"`
	MAType     string `json:"ma_type"`
}

// BbandsData represents individual Bollinger Bands data points.
type BbandsData struct {
	Datetime   string `json:"datetime"`
	UpperBand  string `json:"upper_band"`
	MiddleBand string `json:"middle_band"`
	LowerBand  string `json:"lower_band"`
}
