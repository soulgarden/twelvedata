package request

// GetIPOCalendar represents a request to retrieve IPO calendar data.
// Supports filtering by date range, output size, and pagination.
type GetIPOCalendar struct {
	APIKey
	StartDate  string `schema:"start_date,omitempty"`
	EndDate    string `schema:"end_date,omitempty"`
	OutputSize int    `schema:"outputsize,omitempty"`
	Page       int    `schema:"page,omitempty"`
}
