package response

// Funds represents the response structure for funds data.
type Funds struct {
	Data   []*Fund `json:"data"`
	Count  int     `json:"count"`
	Status string  `json:"status"`
}

// Fund represents a single investment fund with its details and access information.
type Fund struct {
	Symbol   string      `json:"symbol"`
	Name     string      `json:"name"`
	Country  string      `json:"country"`
	Currency string      `json:"currency"`
	Exchange string      `json:"exchange"`
	MicCode  string      `json:"mic_code"`
	Type     string      `json:"type"`
	FigiCode string      `json:"figi_code"`
	Isin     string      `json:"isin"`
	Cusip    string      `json:"cusip"`
	Access   *FundAccess `json:"access"`
}

// FundAccess represents access level information for fund data.
type FundAccess struct {
	Global string `json:"global"`
	Plan   string `json:"plan"`
}
