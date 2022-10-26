package response

type Indices struct {
	Data []Index `json:"data"`
}

type Index struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Currency string `json:"currency"`
}
