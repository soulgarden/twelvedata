package request

// GetAnalystRatingsSnapshot represents request parameters for analyst ratings snapshot.
type GetAnalystRatingsSnapshot struct {
	APIKey
	Symbol       string `schema:"symbol,omitempty"`
	Figi         string `schema:"figi,omitempty"`
	Isin         string `schema:"isin,omitempty"`
	Cusip        string `schema:"cusip,omitempty"`
	Exchange     string `schema:"exchange,omitempty"`
	Country      string `schema:"country,omitempty"`
	RatingChange string `schema:"rating_change,omitempty"`
	OutputSize   int    `schema:"outputsize,omitempty"`
}
