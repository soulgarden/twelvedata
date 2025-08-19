package request

// GetAnalystRatingsUSEquities represents request parameters for analyst ratings US equities.
type GetAnalystRatingsUSEquities struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	Country  string `schema:"country,omitempty"`
	Format   string `schema:"format,omitempty"`
}
