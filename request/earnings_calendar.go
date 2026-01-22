package request

// GetEarningsCalendar represents request parameters for earnings calendar data.
type GetEarningsCalendar struct {
	APIKey
	Exchange      string `schema:"exchange,omitempty"`
	MicCode       string `schema:"mic_code,omitempty"`
	Country       string `schema:"country,omitempty"`
	Format        string `schema:"format,omitempty"`
	Delimiter     string `schema:"delimiter,omitempty"`
	DecimalPlaces int    `schema:"dp,omitempty"`
	StartDate     string `schema:"start_date,omitempty"`
	EndDate       string `schema:"end_date,omitempty"`
}
