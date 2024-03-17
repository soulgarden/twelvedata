package request

type GetExchanges struct {
	instrumentType, name, code, country string
	showPlan                            bool
}
