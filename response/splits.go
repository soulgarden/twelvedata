package response

import "github.com/guregu/null/v6"

// Splits represents the response structure for stock splits data.
type Splits struct {
	Meta   SplitsMeta   `json:"meta"`
	Splits []SplitEvent `json:"splits"`
}

// SplitsMeta contains metadata for splits data.
type SplitsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

// SplitEvent represents a single stock split event with ratio and date information.
type SplitEvent struct {
	Date        string     `json:"date"`
	Description string     `json:"description"`
	Ratio       null.Float `json:"ratio"`
	FromFactor  null.Int   `json:"from_factor"`
	ToFactor    null.Int   `json:"to_factor"`
}
