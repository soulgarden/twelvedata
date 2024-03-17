package request

type GetCashFlow struct {
	symbol, exchange, micCode, country, startDate, endDate, period string
}
