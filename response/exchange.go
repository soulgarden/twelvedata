package response

type Exchanges struct {
	Data []*Exchange `json:"data"`
}

type Exchange struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Country  string `json:"country"`
	Timezone string `json:"timezone"`
}
