package response

type Stocks struct {
	Data []*Stock `json:"data"`
}

type Stock struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Currency string `json:"currency"`
	Exchange string `json:"exchange"`
	Country  string `json:"country"`
	Type     string `json:"type"`
}
