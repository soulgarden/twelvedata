package response

// CCI represents the Commodity Channel Index technical indicator response data.
type CCI struct {
	Meta   CCIMeta    `json:"meta"`
	Values []CCIValue `json:"values"`
	Status string     `json:"status"`
}

// CCIMeta represents the metadata for the Commodity Channel Index technical indicator response.
type CCIMeta struct {
	Symbol           string       `json:"symbol"`
	Interval         string       `json:"interval"`
	Currency         string       `json:"currency"`
	ExchangeTimezone string       `json:"exchange_timezone"`
	Exchange         string       `json:"exchange"`
	MicCode          string       `json:"mic_code"`
	Type             string       `json:"type"`
	Indicator        CCIIndicator `json:"indicator"`
}

// CCIIndicator contains metadata about the Commodity Channel Index indicator configuration.
type CCIIndicator struct {
	Name       string `json:"name"`
	TimePeriod int    `json:"time_period"`
}

// CCIValue represents a single data point in the Commodity Channel Index technical indicator response.
type CCIValue struct {
	Datetime string `json:"datetime"`
	CCI      string `json:"cci"`
}
