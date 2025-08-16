package response

type Commodities struct {
	Data   []*Commodity `json:"data"`
	Status string       `json:"status"`
}

type Commodity struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
