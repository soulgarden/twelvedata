package response

import "github.com/guregu/null/v6"

// MOM represents the response structure for the Momentum technical indicator endpoint.
type MOM struct {
	Meta   MOMMeta    `json:"meta"`
	Values []MOMValue `json:"values"`
	Status string     `json:"status"`
}

// MOMMeta contains metadata information about the Momentum calculation.
type MOMMeta struct {
	Symbol           string       `json:"symbol"`
	Interval         string       `json:"interval"`
	Currency         string       `json:"currency"`
	ExchangeTimezone string       `json:"exchange_timezone"`
	Exchange         string       `json:"exchange"`
	MicCode          string       `json:"mic_code"`
	Type             string       `json:"type"`
	Indicator        MOMIndicator `json:"indicator"`
}

// MOMIndicator contains metadata about the Momentum indicator configuration.
type MOMIndicator struct {
	Name       string   `json:"name"`
	SeriesType string   `json:"series_type"`
	TimePeriod null.Int `json:"time_period"`
}

// MOMValue represents individual Momentum data points.
type MOMValue struct {
	Datetime string     `json:"datetime"`
	MOM      null.Float `json:"mom"`
}
