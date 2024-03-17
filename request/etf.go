package request

type GetEtfs struct {
	symbol, exchange, micCode, country string
	showPlan, includeDelisted          bool
}
