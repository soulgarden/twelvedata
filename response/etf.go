package response

// Etfs represents the response structure for ETF data.
type Etfs struct {
	Data   []Etf  `json:"data"`
	Status string `json:"status"`
}

// Etf represents a single exchange-traded fund with its details and access information.
type Etf struct {
	Symbol   string  `json:"symbol"`
	Name     string  `json:"name"`
	Currency string  `json:"currency"`
	Exchange string  `json:"exchange"`
	MicCode  string  `json:"mic_code"`
	Country  string  `json:"country"`
	FigiCode string  `json:"figi_code"`
	Access   *Access `json:"access"`
}
