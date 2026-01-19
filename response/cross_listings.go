package response

// CrossListings represents the response structure for cross-listing data.
type CrossListings struct {
	Result CrossListingsResult `json:"result"`
	Status string              `json:"status"`
}

// CrossListingsResult contains the cross-listing results.
type CrossListingsResult struct {
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
