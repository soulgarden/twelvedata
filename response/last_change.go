package response

// LastChange represents the response from Last Changes endpoint /last_change/{endpoint}.
type LastChange struct {
	Pagination LastChangePagination `json:"pagination"`
	Data       []LastChangeData     `json:"data"`
}

// LastChangePagination contains pagination information for Last Changes response.
type LastChangePagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
}

// LastChangeData represents individual change record in the Last Changes response.
type LastChangeData struct {
	Symbol      string `json:"symbol"`
	Exchange    string `json:"exchange"`
	Country     string `json:"country"`
	Endpoint    string `json:"endpoint"`
	LastChange  string `json:"last_change"`
	ChangeType  string `json:"change_type"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}
