package twelvedata

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/gorilla/schema"
	"github.com/soulgarden/twelvedata/response"
	"github.com/valyala/fasthttp"
)

var encoder = schema.NewEncoder()

// Endpoint represents a generic HTTP endpoint with type-safe request/response handling.
type Endpoint[Request any, Response any, Credits response.Credits, Error error] struct {
	httpCli *HTTPCli
	URL     string
}

// NewEndpoint creates a new endpoint instance with the specified HTTP client and URI.
func NewEndpoint[Request any, Response any, Credits response.Credits, Error error](httpCli *HTTPCli, uri string) *Endpoint[Request, Response, Credits, Error] {
	return &Endpoint[Request, Response, Credits, Error]{
		httpCli: httpCli,
		URL:     uri,
	}
}

// Call executes the endpoint request and returns the response, credits, and any errors.
func (endpoint Endpoint[Request, Response, Credits, ErrorResponse]) Call(req Request) (resp Response, creds response.Credits, err Error) {
	httpResp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(httpResp)

	var (
		creditsLeft, creditsUsed int64
		innerErr                 error
	)

	values := url.Values{}

	if innerErr = encoder.Encode(req, values); innerErr != nil {
		return resp, creds, NewError[Error](fmt.Errorf("encoding url params: %w", innerErr), nil)
	}

	var uri *url.URL

	if uri, innerErr = url.Parse(endpoint.URL); innerErr != nil {
		return resp, creds, NewError[Error](fmt.Errorf("parse uri: %w", innerErr), nil)
	}

	uri.RawQuery = values.Encode()

	if creditsLeft, creditsUsed, innerErr = endpoint.httpCli.makeRequest(uri.String(), httpResp); innerErr != nil {
		// Check if it's a network or timeout error
		if isTimeoutError(innerErr) {
			return resp, creds, NewError[Error](&TimeoutError{Message: innerErr.Error()}, nil)
		}

		if isNetworkError(innerErr) {
			return resp, creds, NewError[Error](&NetworkError{Message: innerErr.Error(), Cause: innerErr}, nil)
		}

		return resp, creds, NewError[Error](innerErr, nil)
	}

	creds = &response.CreditsImpl{}

	creds.SetCreditsLeft(creditsLeft)
	creds.SetCreditsUsed(creditsUsed)

	// Handle HTTP status code errors first
	statusCode := httpResp.StatusCode()
	if statusCode >= 400 {
		var apiError response.Error

		var parsedAPIError *response.Error

		// Try to parse API error from response body
		if innerErr := json.Unmarshal(httpResp.Body(), &apiError); innerErr == nil && apiError.Status == "error" {
			parsedAPIError = &apiError
		}

		// Check for domain-specific errors first
		if parsedAPIError != nil {
			if domainErr := ParseDomainError(parsedAPIError, statusCode, uri.String()); domainErr != nil {
				return resp, creds, NewError[Error](domainErr, nil)
			}
		}

		// Fall back to HTTP error types
		typedErr := NewHTTPError(statusCode, httpResp.Body(), uri.String(), parsedAPIError, nil)

		return resp, creds, NewError[Error](typedErr, nil)
	}

	var respErr response.Error

	if innerErr := json.Unmarshal(httpResp.Body(), &respErr); innerErr == nil && respErr.Status == "error" {
		// Check for domain-specific errors in 200 OK responses with error status
		if domainErr := ParseDomainError(&respErr, statusCode, uri.String()); domainErr != nil {
			return resp, creds, NewError[Error](domainErr, nil)
		}

		// Fall back to generic error
		return resp, creds, NewError[Error](fmt.Errorf("error received: %s", respErr.Error()), respErr)
	}

	if innerErr := json.Unmarshal(httpResp.Body(), &resp); innerErr != nil {
		return resp, creds, NewError[Error](fmt.Errorf("unmarshall json: %w", innerErr), nil)
	}

	return resp, creds, err
}

// NewError creates a new generic error wrapper.
func NewError[T error](err error, t T) ErrImplError[T] {
	return ErrImplError[T]{
		generic: t,
		inner:   err,
	}
}

// Error represents a generic error interface.
type Error interface {
	Error() string
}

// ErrImplError represents a generic error implementation with type safety.
type ErrImplError[Err error] struct {
	generic Err
	inner   error
}

func (e ErrImplError[Err]) Error() string {
	return e.inner.Error()
}

func (e ErrImplError[Err]) Unwrap() error {
	return e.inner
}

// Helper functions to identify error types.
func isTimeoutError(err error) bool {
	// Check for common timeout error patterns
	errStr := err.Error()

	return strings.Contains(errStr, "timeout") ||
		strings.Contains(errStr, "deadline exceeded") ||
		strings.Contains(errStr, "context deadline exceeded")
}

func isNetworkError(err error) bool {
	// Check for common network error patterns
	errStr := err.Error()

	return strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "no route to host") ||
		strings.Contains(errStr, "network is unreachable") ||
		strings.Contains(errStr, "connection reset")
}
