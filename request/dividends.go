package request

type GetDividends struct {
	symbol, exchange, micCode, country, r, startDate, endDate string
}
