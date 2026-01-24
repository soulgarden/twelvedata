package response

import (
	"encoding/json"
	"github.com/guregu/null/v6"
)

// Batches represents the response from the batch API endpoint.
type Batches struct {
	Code   null.Int                 `json:"code"`
	Status string                   `json:"status"`
	Data   map[string]BatchResponse `json:"data"`
}

// BatchResponse represents a single response within a batch operation.
type BatchResponse struct {
	Status   string          `json:"status"`
	Response json.RawMessage `json:"response"`
}
