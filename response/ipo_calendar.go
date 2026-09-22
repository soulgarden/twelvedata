package response

import "github.com/guregu/null/v6"

// IPOCalendar groups IPO entries by date (YYYY-MM-DD).
type IPOCalendar map[string][]IPOCalendarData

// IPOCalendarData describes an offering. The date is the containing map key.
type IPOCalendarData struct {
	Symbol         string     `json:"symbol"`
	Name           string     `json:"name"`
	Exchange       string     `json:"exchange"`
	MicCode        string     `json:"mic_code"`
	PriceRangeLow  null.Float `json:"price_range_low"`
	PriceRangeHigh null.Float `json:"price_range_high"`
	OfferPrice     null.Float `json:"offer_price"`
	Currency       string     `json:"currency"`
	Shares         null.Int   `json:"shares"`
}
