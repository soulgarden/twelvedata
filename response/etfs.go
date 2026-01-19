package response

// ETFs represents the response structure for ETFs catalog data.
type ETFs struct {
	Data   []*ETF `json:"data"`
	Status string `json:"status"`
}

// ETF represents a single ETF instrument with its details and access information.
type ETF struct {
	Symbol   string     `json:"symbol"`
	Name     string     `json:"name"`
	Currency string     `json:"currency"`
	Exchange string     `json:"exchange"`
	MicCode  string     `json:"mic_code"`
	Country  string     `json:"country"`
	FigiCode string     `json:"figi_code"`
	CfiCode  string     `json:"cfi_code"`
	Isin     string     `json:"isin"`
	Cusip    string     `json:"cusip"`
	Access   *ETFAccess `json:"access"`
}

// ETFAccess represents access information for ETF data.
type ETFAccess struct {
	Global string `json:"global"`
	Plan   string `json:"plan"`
}
