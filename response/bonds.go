package response

type Bonds struct {
	Data   []*Bond `json:"data"`
	Count  int     `json:"count"`
	Status string  `json:"status"`
}

type Bond struct {
	Symbol   string      `json:"symbol"`
	Name     string      `json:"name"`
	Country  string      `json:"country"`
	Currency string      `json:"currency"`
	Exchange string      `json:"exchange"`
	MicCode  string      `json:"mic_code"`
	Type     string      `json:"type"`
	Access   *BondAccess `json:"access"`
}

type BondAccess struct {
	Global string `json:"global"`
	Plan   string `json:"plan"`
}
