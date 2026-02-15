package request

// GetCrossListings represents request parameters for cross listings data.
type GetCrossListings struct {
	APIKey
	Symbol   string `schema:"symbol"`
	Exchange string `schema:"exchange,omitempty"`
	MicCode  string `schema:"mic_code,omitempty"`
	Country  string `schema:"country,omitempty"`
}
