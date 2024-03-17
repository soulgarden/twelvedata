package request

type GetIncomeStatement struct {
	symbol, exchange, micCode, country, period, startDate, endDate string
}
