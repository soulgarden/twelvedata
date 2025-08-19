package response

// CrossListings represents the response structure for cross-listing data.
type CrossListings struct {
	Count int             `json:"count"`
	List  []*CrossListing `json:"list"`
}

// CrossListing represents a single cross-listed security across exchanges.
type CrossListing struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
	MicCode  string `json:"mic_code"`
}
