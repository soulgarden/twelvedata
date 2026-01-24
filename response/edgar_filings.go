package response

import "github.com/guregu/null/v6"

// EDGARFilings represents the response structure for EDGAR filings data.
type EDGARFilings struct {
	Meta   EDGARFilingsMeta `json:"meta"`
	Values []EDGARFiling    `json:"values"`
}

// EDGARFilingsMeta contains metadata for EDGAR filings data.
type EDGARFilingsMeta struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	MicCode  string `json:"mic_code"`
	Type     string `json:"type"`
}

// EDGARFiling represents a single EDGAR filing record.
type EDGARFiling struct {
	Cik       null.Int          `json:"cik"`
	FiledAt   null.Int          `json:"filed_at"`
	Files     []EDGARFilingFile `json:"files"`
	FilingURL string            `json:"filing_url"`
	FormType  string            `json:"form_type"`
	Ticker    []string          `json:"ticker"`
}

// EDGARFilingFile represents a single EDGAR filing file.
type EDGARFilingFile struct {
	Name string   `json:"name"`
	Size null.Int `json:"size"`
	Type string   `json:"type"`
	URL  string   `json:"url"`
}
