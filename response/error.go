package response

import "fmt"

// Error represents an API error response with code, message and status.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s, status: %s", e.Code, e.Message, e.Status)
}
