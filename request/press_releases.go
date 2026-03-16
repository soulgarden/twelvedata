package request

// GetPressReleases represents request parameters for press releases data.
// One of Symbol, Figi, Isin, or Cusip is required by the API.
type GetPressReleases struct {
	APIKey
	Symbol     string `schema:"symbol,omitempty"`
	Figi       string `schema:"figi,omitempty"`
	Isin       string `schema:"isin,omitempty"`
	Cusip      string `schema:"cusip,omitempty"`
	Exchange   string `schema:"exchange,omitempty"`
	MicCode    string `schema:"mic_code,omitempty"`
	StartDate  string `schema:"start_date,omitempty"`
	EndDate    string `schema:"end_date,omitempty"`
	Language   string `schema:"language,omitempty"`
	TimeZone   string `schema:"timezone,omitempty"`
	OutputSize int    `schema:"outputsize,omitempty"`
}
