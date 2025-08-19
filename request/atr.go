package request

// GetATR represents the request parameters for the Average True Range (ATR) technical indicator endpoint.
type GetATR struct {
	APIKey
	Symbol        string `schema:"symbol,omitempty"`
	Interval      string `schema:"interval,omitempty"`
	Exchange      string `schema:"exchange,omitempty"`
	MICCode       string `schema:"mic_code,omitempty"`
	Country       string `schema:"country,omitempty"`
	TimePeriod    int    `schema:"time_period,omitempty"`
	SeriesType    string `schema:"series_type,omitempty"`
	OutputSize    int    `schema:"outputsize,omitempty"`
	Format        string `schema:"format,omitempty"`
	Delimiter     string `schema:"delimiter,omitempty"`
	Prepost       bool   `schema:"prepost,omitempty"`
	DP            int    `schema:"dp,omitempty"`
	Order         string `schema:"order,omitempty"`
	StartDate     string `schema:"start_date,omitempty"`
	EndDate       string `schema:"end_date,omitempty"`
	PreviousClose bool   `schema:"previous_close,omitempty"`
}
