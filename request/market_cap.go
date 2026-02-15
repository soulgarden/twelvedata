package request

// GetMarketCap represents request parameters for market capitalization data.
type GetMarketCap struct {
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
	Page       int    `schema:"page,omitempty"`
	OutputSize int    `schema:"outputsize,omitempty"`
}
