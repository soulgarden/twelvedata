package request

type GetEarningsCalendar struct {
	ApiKey
	Exchange      string `schema:"exchange,omitempty"`
	MicCode       string `schema:"mic_code,omitempty"`
	Country       string `schema:"country,omitempty"`
	DecimalPlaces int    `schema:"dp,omitempty"`
	StartDate     string `schema:"start_date,omitempty"`
	EndDate       string `schema:"end_date,omitempty"`
}
