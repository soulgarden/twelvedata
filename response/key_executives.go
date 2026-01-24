package response

import "github.com/guregu/null/v6"

// KeyExecutives represents the response structure for the Key Executives endpoint.
// Contains metadata about the company and a list of key executives with their details.
type KeyExecutives struct {
	Meta          KeyExecutivesMeta `json:"meta"`
	KeyExecutives []KeyExecutive    `json:"key_executives"`
}

// KeyExecutivesMeta contains metadata information about the company.
type KeyExecutivesMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

// KeyExecutive represents an individual executive's information.
// Pay is nullable as some executives may not have compensation data disclosed.
// Age and YearBorn can be 0 when information is not available.
type KeyExecutive struct {
	Name     string   `json:"name"`
	Title    string   `json:"title"`
	Age      null.Int `json:"age"`
	YearBorn null.Int `json:"year_born"`
	Pay      null.Int `json:"pay"`
}
