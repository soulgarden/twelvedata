package request

type GetEarningsCalendar struct {
	exchange, micCode, country string
	decimalPlaces              int
	startDate, endDate         string
}
