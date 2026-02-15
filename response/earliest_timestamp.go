package response

import "github.com/guregu/null/v6"

// EarliestTimestamp represents the earliest available data timestamp for a symbol.
type EarliestTimestamp struct {
	Datetime string   `json:"datetime"`
	UnixTime null.Int `json:"unix_time"`
}
