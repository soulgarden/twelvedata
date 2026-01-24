package response

import "github.com/guregu/null/v6"

// CrossListings represents the response structure for cross-listing data.
type CrossListings struct {
	Result CrossListingsResult `json:"result"`
	Status string              `json:"status"`
}

// CrossListingsResult contains the cross-listing results.
type CrossListingsResult struct {
	Count null.Int        `json:"count"`
	List  []*CrossListing `json:"list"`
}

// CrossListing represents a single cross-listed security across exchanges.
type CrossListing struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
	MicCode  string `json:"mic_code"`
}
