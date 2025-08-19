package response

// Bonds represents the response structure for bonds data from the API.
type Bonds struct {
	Data   []*Bond `json:"data"`
	Count  int     `json:"count"`
	Status string  `json:"status"`
}

// Bond represents a single bond instrument with its details and access information.
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

// BondAccess represents access level information for bond data.
type BondAccess struct {
	Global string `json:"global"`
	Plan   string `json:"plan"`
}
