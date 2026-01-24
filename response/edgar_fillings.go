package response

import "github.com/guregu/null/v6"

// EDGARFillings represents the response structure for EDGAR filings data.
type EDGARFillings struct {
	Meta   EDGARFillingsMeta `json:"meta"`
	Values []EDGARFilling    `json:"values"`
}

// EDGARFillingsMeta contains metadata for EDGAR filings data.
type EDGARFillingsMeta struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	MicCode  string `json:"mic_code"`
	Type     string `json:"type"`
}

// EDGARFilling represents a single EDGAR filing record.
type EDGARFilling struct {
	Cik       null.Int           `json:"cik"`
	FiledAt   null.Int           `json:"filed_at"`
	Files     []EDGARFillingFile `json:"files"`
	FilingURL string             `json:"filing_url"`
	FormType  string             `json:"form_type"`
	Ticker    []string           `json:"ticker"`
}

// EDGARFillingFile represents a single EDGAR filing file.
type EDGARFillingFile struct {
	Name string   `json:"name"`
	Size null.Int `json:"size"`
	Type string   `json:"type"`
	URL  string   `json:"url"`
}
