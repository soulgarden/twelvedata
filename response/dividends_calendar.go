package response

import "github.com/guregu/null/v6"

// DividendsCalendar is a slice of dividend calendar events.
type DividendsCalendar []DividendCalendarEvent

// DividendCalendarEvent represents a single dividend calendar event with timing and amount details.
type DividendCalendarEvent struct {
	Symbol   string     `json:"symbol"`
	MicCode  string     `json:"mic_code"`
	Exchange string     `json:"exchange"`
	ExDate   string     `json:"ex_date"`
	Amount   null.Float `json:"amount"`
}
