package response

import "github.com/guregu/null/v6"

// Stocks represents the response structure for stock data.
type Stocks struct {
	Data   []*Stock `json:"data"`
	Count  null.Int `json:"count"`
	Status string   `json:"status"`
}

// Stock represents a single stock instrument with its details and access information.
type Stock struct {
	Symbol   string       `json:"symbol"`
	Name     string       `json:"name"`
	Currency string       `json:"currency"`
	Exchange string       `json:"exchange"`
	MicCode  string       `json:"mic_code"`
	Country  string       `json:"country"`
	Type     string       `json:"type"`
	FigiCode string       `json:"figi_code"`
	CfiCode  string       `json:"cfi_code"`
	Isin     string       `json:"isin"`
	Cusip    string       `json:"cusip"`
	Access   *StockAccess `json:"access"`
}

// StockAccess represents access level information for stock data.
type StockAccess struct {
	Global string `json:"global"`
	Plan   string `json:"plan"`
}
