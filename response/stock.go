package response

type Stocks struct {
	Data   []*Stock `json:"data"`
	Status string   `json:"status"`
}

type Stock struct {
	Symbol   string       `json:"symbol"`
	Name     string       `json:"name"`
	Currency string       `json:"currency"`
	Exchange string       `json:"exchange"`
	MicCode  string       `json:"mic_code"`
	Country  string       `json:"country"`
	Type     string       `json:"type"`
	FigiCode string       `json:"figi_code"`
	Isin     string       `json:"isin"`
	Access   *StockAccess `json:"access"`
}

type StockAccess struct {
	Global string `json:"global"`
	Plan   string `json:"plan"`
}
