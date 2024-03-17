package request

type GetBalanceSheet struct {
	symbol, exchange, micCode, country, period, startDate, endDate string
}
