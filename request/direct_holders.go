package request

// GetDirectHolders represents request parameters for direct holders data.
type GetDirectHolders struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	MicCode  string `schema:"mic_code,omitempty"`
	Country  string `schema:"country,omitempty"`
	Format   string `schema:"format,omitempty"`
}
