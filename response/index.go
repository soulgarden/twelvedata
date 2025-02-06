package response

type Indices struct {
	Data   []Index `json:"data"`
	Count  int     `json:"count"`
	Status string  `json:"status"`
}

type Index struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Currency string `json:"currency"`
	Exchange string `json:"exchange"`
	MicCode  string `json:"mic_code"`
}
