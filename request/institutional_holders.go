package request

// GetInstitutionalHolders represents request parameters for institutional holders data.
type GetInstitutionalHolders struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	MicCode  string `schema:"mic_code,omitempty"`
	Country  string `schema:"country,omitempty"`
	Format   string `schema:"format,omitempty"`
}
