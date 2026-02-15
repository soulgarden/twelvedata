package request

// GetUsage represents request parameters for usage data.
type GetUsage struct {
	APIKey
	Format    string `schema:"format,omitempty"`
	Delimiter string `schema:"delimiter,omitempty"`
	TimeZone  string `schema:"timezone,omitempty"`
}
