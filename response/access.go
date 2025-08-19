// Package response contains all response structures for Twelve Data API endpoints.
// These structures represent the JSON responses returned by various API calls
// and are used throughout the client library for type-safe API interactions.
package response

// Access represents access level information for API resources.
type Access struct {
	Global string `json:"global"`
	Plan   string `json:"plan"`
}
