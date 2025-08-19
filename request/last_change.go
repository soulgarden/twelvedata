package request

// GetLastChange represents request for Last Changes endpoint /last_change/{endpoint}
// The endpoint parameter is substituted in the URL path.
type GetLastChange struct {
	APIKey
	Endpoint  string `schema:"-"`                    // URL path parameter, not query param
	Symbol    string `schema:"symbol,omitempty"`     // Symbol to track changes for
	Exchange  string `schema:"exchange,omitempty"`   // Exchange for the symbol
	Country   string `schema:"country,omitempty"`    // Country filter
	Date      string `schema:"date,omitempty"`       // Specific date to check changes
	StartDate string `schema:"start_date,omitempty"` // Start date range
	EndDate   string `schema:"end_date,omitempty"`   // End date range
	Page      int    `schema:"page,omitempty"`       // Pagination page number
	PerPage   int    `schema:"per_page,omitempty"`   // Items per page
}
