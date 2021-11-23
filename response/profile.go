package response

// nolint: tagliatelle
type Profile struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Exchange    string `json:"exchange"`
	Sector      string `json:"sector"`
	Industry    string `json:"industry"`
	Employees   int    `json:"employees"`
	Website     string `json:"website"`
	Description string `json:"description"`
	Type        string `json:"type"`
	CEO         string `json:"CEO"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Zip         string `json:"zip"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
}
