package request

// GetMarketCap represents request parameters for market capitalization data.
type GetMarketCap struct {
	APIKey
	Symbol     string `schema:"symbol,omitempty"`
	Exchange   string `schema:"exchange,omitempty"`
	MicCode    string `schema:"mic_code,omitempty"`
	Country    string `schema:"country,omitempty"`
	Interval   string `schema:"interval,omitempty"`
	OutputSize string `schema:"outputsize,omitempty"`
	Format     string `schema:"format,omitempty"`
	Delimiter  string `schema:"delimiter,omitempty"`
	StartDate  string `schema:"start_date,omitempty"`
	EndDate    string `schema:"end_date,omitempty"`
}
