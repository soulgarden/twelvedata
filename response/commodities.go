package response

// Commodities represents the response structure for commodities data.
type Commodities struct {
	Data   []*Commodity `json:"data"`
	Status string       `json:"status"`
}

// Commodity represents a single commodity instrument with its details.
type Commodity struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
