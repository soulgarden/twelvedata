package response

import "github.com/guregu/null/v6"

// DirectHolders represents the response structure for direct holders data.
type DirectHolders struct {
	Meta          DirectHoldersMeta `json:"meta"`
	DirectHolders []DirectHolder    `json:"direct_holders"`
}

// DirectHoldersMeta contains metadata for direct holders data.
type DirectHoldersMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

// DirectHolder represents a single direct holder record.
type DirectHolder struct {
	HolderName   string      `json:"holder_name"`
	DateReported string      `json:"date_reported"`
	Shares       int64       `json:"shares"`
	Value        null.Int    `json:"value"`
	PercentHeld  null.Float  `json:"percent_held"`
	HolderType   null.String `json:"holder_type"`
}
