package request

type GetInsiderTransactions struct {
	ApiKey
	Symbol   string `schema:"symbol"`
	Exchange string `schema:"exchange"`
	MicCode  string `schema:"mic_code"`
	Country  string `schema:"country"`
}
