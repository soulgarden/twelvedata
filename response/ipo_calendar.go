package response

// IPOCalendar represents a collection of IPO calendar entries.
type IPOCalendar []IPOCalendarData

// IPOCalendarData represents a single IPO calendar entry containing
// company information and the IPO date.
type IPOCalendarData struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
	Date     string `json:"date"`
}
