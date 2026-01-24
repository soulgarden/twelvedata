package request

import "net/url"

// GetBatches represents request parameters for batch operations.
// The batch endpoint accepts a JSON POST request with multiple API calls.
type GetBatches struct {
	APIKey
	// Requests is a map of request IDs to batch request items.
	Requests map[string]BatchRequest
}

// BatchRequest represents a single request within a batch operation.
type BatchRequest struct {
	// URL is the API endpoint URL with parameters (e.g., "/time_series?symbol=AAPL&interval=1min&apikey=demo")
	URL string `json:"url"`
}

// Method returns the HTTP method for batch requests.
func (b GetBatches) Method() string {
	return "POST"
}

// Headers returns the headers required for batch requests.
func (b GetBatches) Headers() map[string]string {
	if b.APIKey.APIKey == "" {
		return nil
	}

	return map[string]string{
		"Authorization": "apikey " + b.APIKey.APIKey,
	}
}

// Body returns the JSON body for batch requests.
func (b GetBatches) Body() (any, string, error) {
	if b.Requests == nil {
		return map[string]BatchRequest{}, "application/json", nil
	}

	return b.Requests, "application/json", nil
}

// Query returns an empty query set because batch requests use JSON body.
func (b GetBatches) Query() (url.Values, error) {
	return url.Values{}, nil
}
