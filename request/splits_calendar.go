package request

// GetSplitsCalendar represents request parameters for splits calendar data.
type GetSplitsCalendar struct {
	APIKey
	Symbol     string `schema:"symbol,omitempty"`
	Figi       string `schema:"figi,omitempty"`
	Isin       string `schema:"isin,omitempty"`
	Cusip      string `schema:"cusip,omitempty"`
	Exchange   string `schema:"exchange,omitempty"`
	MicCode    string `schema:"mic_code,omitempty"`
	Country    string `schema:"country,omitempty"`
	StartDate  string `schema:"start_date,omitempty"`
	EndDate    string `schema:"end_date,omitempty"`
	OutputSize int    `schema:"outputsize,omitempty"`
	Page       int    `schema:"page,omitempty"`
}
