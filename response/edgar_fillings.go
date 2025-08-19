package response

import "github.com/guregu/null/v6"

// EDGARFillings represents the response structure for EDGAR filings data.
type EDGARFillings struct {
	Meta     EDGARFillingsMeta `json:"meta"`
	Fillings []EDGARFilling    `json:"fillings"`
}

// EDGARFillingsMeta contains metadata for EDGAR filings data.
type EDGARFillingsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

// EDGARFilling represents a single EDGAR filing document.
type EDGARFilling struct {
	FormType        string      `json:"form_type"`
	FilingDate      string      `json:"filing_date"`
	AcceptedDate    string      `json:"accepted_date"`
	ReportDate      null.String `json:"report_date"`
	FilingURL       string      `json:"filing_url"`
	FilingNumber    string      `json:"filing_number"`
	Items           null.String `json:"items"`
	Size            null.Int    `json:"size"`
	IsXBRL          null.Bool   `json:"is_xbrl"`
	IsInlineXBRL    null.Bool   `json:"is_inline_xbrl"`
	PrimaryDocument string      `json:"primary_document"`
	Description     null.String `json:"description"`
}
