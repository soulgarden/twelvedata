package request

// GetBatches represents request parameters for batch operations.
// The batch endpoint accepts a JSON POST request with multiple API calls.
type GetBatches struct {
	APIKey
	// Requests is a map of request IDs to batch request items
	Requests map[string]BatchRequest `json:",inline"`
}

// BatchRequest represents a single request within a batch operation.
type BatchRequest struct {
	// URL is the API endpoint URL with parameters (e.g., "/time_series?symbol=AAPL&interval=1min&apikey=demo")
	URL string `json:"url"`
}

// NewBatchRequest creates a new batch request with the given request map.
func NewBatchRequest(requests map[string]BatchRequest) GetBatches {
	return GetBatches{
		Requests: requests,
	}
}

// AddRequest adds a new request to the batch with the given ID and URL.
func (b *GetBatches) AddRequest(id, url string) {
	if b.Requests == nil {
		b.Requests = make(map[string]BatchRequest)
	}
	b.Requests[id] = BatchRequest{URL: url}
}
