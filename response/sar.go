package response

// SAR represents the Parabolic SAR technical indicator response data.
type SAR struct {
	Meta   SARMeta    `json:"meta"`
	Values []SARValue `json:"values"`
	Status string     `json:"status"`
}

// SARMeta represents the metadata for the Parabolic SAR technical indicator response.
type SARMeta struct {
	Symbol           string       `json:"symbol"`
	Interval         string       `json:"interval"`
	Currency         string       `json:"currency"`
	ExchangeTimezone string       `json:"exchange_timezone"`
	Exchange         string       `json:"exchange"`
	MicCode          string       `json:"mic_code"`
	Type             string       `json:"type"`
	Indicator        SARIndicator `json:"indicator"`
}

// SARIndicator contains metadata about the Parabolic SAR indicator configuration.
type SARIndicator struct {
	Name         string  `json:"name"`
	Acceleration float64 `json:"acceleration"`
	Maximum      float64 `json:"maximum"`
}

// SARValue represents a single data point in the Parabolic SAR technical indicator response.
type SARValue struct {
	Datetime string `json:"datetime"`
	SAR      string `json:"sar"`
}
