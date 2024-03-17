package request

type GetStock struct {
	ApiKey
	symbol          string `schema:"symbol"`
	exchange        string `schema:"exchange"`
	micCode         string `schema:"mic_code"`
	country         string `schema:"country"`
	instrumentType  string `schema:"type"`
	showPlan        bool   `schema:"show_plan"`
	includeDelisted bool   `schema:"include_delisted"`
}
