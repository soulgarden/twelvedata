package request

// GetInsiderTransactions represents request parameters for insider transactions data.
type GetInsiderTransactions struct {
	APIKey
	Symbol   string `schema:"symbol"`
	Exchange string `schema:"exchange"`
	MicCode  string `schema:"mic_code"`
	Country  string `schema:"country"`
}
