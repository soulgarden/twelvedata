// Package request contains structures for API request parameters
package request

// APIKey represents the API key authentication structure.
type APIKey struct {
	APIKey string `schema:"apikey,omitempty"`
}
