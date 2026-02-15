package response

// Countries represents the response structure for country data.
type Countries struct {
	Data   []*Country `json:"data"`
	Status string     `json:"status"`
}

// Country represents a single country with its details and currency information.
type Country struct {
	Iso2         string `json:"iso2"`
	Iso3         string `json:"iso3"`
	Numeric      string `json:"numeric"`
	Name         string `json:"name"`
	OfficialName string `json:"official_name"`
	Capital      string `json:"capital"`
	Currency     string `json:"currency"`
}
