package response

type Etfs struct {
	Data []Etf `json:"data"`
}

type Etf struct {
	Symbol   string  `json:"symbol"`
	Name     string  `json:"name"`
	Currency string  `json:"currency"`
	Exchange string  `json:"exchange"`
	MicCode  string  `json:"mic_code"`
	Country  string  `json:"country"`
	Access   *Access `json:"access"`
}
