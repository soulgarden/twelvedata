package response

import "github.com/guregu/null/v6"

// VWAP represents the Volume Weighted Average Price technical indicator response data.
type VWAP struct {
	Meta   VWAPMeta    `json:"meta"`
	Values []VWAPValue `json:"values"`
	Status string      `json:"status"`
}

// VWAPMeta represents the metadata for the Volume Weighted Average Price technical indicator response.
type VWAPMeta struct {
	Symbol           string        `json:"symbol"`
	Interval         string        `json:"interval"`
	Currency         string        `json:"currency"`
	ExchangeTimezone string        `json:"exchange_timezone"`
	Exchange         string        `json:"exchange"`
	MicCode          string        `json:"mic_code"`
	Type             string        `json:"type"`
	Indicator        VWAPIndicator `json:"indicator"`
}

// VWAPIndicator contains metadata about the Volume Weighted Average Price indicator configuration.
type VWAPIndicator struct {
	Name         string      `json:"name"`
	SDTimePeriod null.Int    `json:"sd_time_period"`
	SD           FloatString `json:"sd"`
}

// VWAPValue represents a single data point in the Volume Weighted Average Price technical indicator response.
type VWAPValue struct {
	Datetime  string      `json:"datetime"`
	VWAPLower FloatString `json:"vwap_lower"`
	VWAP      FloatString `json:"vwap"`
	VWAPUpper FloatString `json:"vwap_upper"`
}
