package request

type GetExchangeRate struct {
	symbol, date, timeZone string
	decimalPlaces          int
}
