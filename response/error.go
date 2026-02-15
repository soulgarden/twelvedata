package response

import (
	"fmt"

	"github.com/guregu/null/v6"
)

// Error represents an API error response with code, message and status.
type Error struct {
	Code    null.Int `json:"code"`
	Message string   `json:"message"`
	Status  string   `json:"status"`
}

func (e Error) Error() string {
	code := int64(0)
	if e.Code.Valid {
		code = e.Code.Int64
	}
	return fmt.Sprintf("code: %d, message: %s, status: %s", code, e.Message, e.Status)
}
