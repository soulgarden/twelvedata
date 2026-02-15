package request

// GetDividendsCalendar represents request parameters for dividends calendar data.
type GetDividendsCalendar struct {
	APIKey
	Symbol     string `schema:"symbol,omitempty"`
	FIGI       string `schema:"figi,omitempty"`
	ISIN       string `schema:"isin,omitempty"`
	CUSIP      string `schema:"cusip,omitempty"`
	Exchange   string `schema:"exchange,omitempty"`
	MicCode    string `schema:"mic_code,omitempty"`
	Country    string `schema:"country,omitempty"`
	StartDate  string `schema:"start_date,omitempty"`
	EndDate    string `schema:"end_date,omitempty"`
	OutputSize int    `schema:"outputsize,omitempty"`
	Page       int    `schema:"page,omitempty"`
}
