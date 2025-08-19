package request

// GetAnalystRatingsSnapshot represents request parameters for analyst ratings snapshot.
type GetAnalystRatingsSnapshot struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	Country  string `schema:"country,omitempty"`
	Format   string `schema:"format,omitempty"`
}
