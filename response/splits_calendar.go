package response

import "github.com/guregu/null/v6"

// SplitsCalendar is a slice of stock split calendar events.
type SplitsCalendar []SplitsCalendarItem

// SplitsCalendarItem represents a single stock split calendar event.
type SplitsCalendarItem struct {
	Date        string     `json:"date"`
	Symbol      string     `json:"symbol"`
	MicCode     string     `json:"mic_code"`
	Exchange    string     `json:"exchange"`
	Description string     `json:"description"`
	Ratio       null.Float `json:"ratio"`
	FromFactor  null.Int   `json:"from_factor"`
	ToFactor    null.Int   `json:"to_factor"`
}
