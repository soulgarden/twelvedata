package response

// Exchanges represents the response structure for exchange data.
type Exchanges struct {
	Data   []Exchange `json:"data"`
	Status string     `json:"status"`
}

// Exchange represents a single financial exchange with its details and access information.
type Exchange struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Country  string `json:"country"`
	Timezone string `json:"timezone"`

	Access *Access `json:"access"`
}
