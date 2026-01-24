package response

import "github.com/guregu/null/v6"

// InstitutionalHolders represents the response structure for institutional holders data.
type InstitutionalHolders struct {
	Meta                 InstitutionalHoldersMeta `json:"meta"`
	InstitutionalHolders []InstitutionalHolder    `json:"institutional_holders"`
}

// InstitutionalHoldersMeta contains metadata for institutional holders data.
type InstitutionalHoldersMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

// InstitutionalHolder represents a single institutional holder record.
type InstitutionalHolder struct {
	EntityName   string     `json:"entity_name"`
	DateReported string     `json:"date_reported"`
	Shares       null.Int   `json:"shares"`
	Value        null.Int   `json:"value"`
	PercentHeld  null.Float `json:"percent_held"`
}
