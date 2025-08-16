package response

type CrossListings struct {
	Count int             `json:"count"`
	List  []*CrossListing `json:"list"`
}

type CrossListing struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
	MicCode  string `json:"mic_code"`
}
