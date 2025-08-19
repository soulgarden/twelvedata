package response

// SplitsCalendar is a slice of stock split calendar events.
type SplitsCalendar []SplitsCalendarItem

// SplitsCalendarItem represents a single stock split calendar event.
type SplitsCalendarItem struct {
	Date        string  `json:"date"`
	Symbol      string  `json:"symbol"`
	MicCode     string  `json:"mic_code"`
	Exchange    string  `json:"exchange"`
	Description string  `json:"description"`
	Ratio       float64 `json:"ratio"`
	FromFactor  int     `json:"from_factor"`
	ToFactor    int     `json:"to_factor"`
}
