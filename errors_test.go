package twelvedata

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/soulgarden/twelvedata/response"
)

func TestNewHTTPError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       []byte
		apiError   *response.Error
		wantType   string
	}{
		{
			name:       "bad request with api error",
			statusCode: http.StatusBadRequest,
			body:       []byte(`{"code":400,"message":"Invalid parameter","status":"error"}`),
			apiError: &response.Error{
				Code:    400,
				Message: "Invalid parameter",
				Status:  "error",
			},
			wantType: "*twelvedata.BadRequestError",
		},
		{
			name:       "unauthorized error",
			statusCode: http.StatusUnauthorized,
			body:       []byte(`{"code":401,"message":"Invalid API key","status":"error"}`),
			apiError: &response.Error{
				Code:    401,
				Message: "Invalid API key",
				Status:  "error",
			},
			wantType: "*twelvedata.UnauthorizedError",
		},
		{
			name:       "not found error",
			statusCode: http.StatusNotFound,
			body:       []byte(`Not Found`),
			apiError:   nil,
			wantType:   "*twelvedata.NotFoundError",
		},
		{
			name:       "rate limit error",
			statusCode: http.StatusTooManyRequests,
			body:       []byte(`{"code":429,"message":"Rate limit exceeded","status":"error"}`),
			apiError: &response.Error{
				Code:    429,
				Message: "Rate limit exceeded",
				Status:  "error",
			},
			wantType: "*twelvedata.TooManyRequestsError",
		},
		{
			name:       "internal server error",
			statusCode: http.StatusInternalServerError,
			body:       []byte(`Internal Server Error`),
			apiError:   nil,
			wantType:   "*twelvedata.InternalServerError",
		},
		{
			name:       "generic http error",
			statusCode: http.StatusBadGateway,
			body:       []byte(`Bad Gateway`),
			apiError:   nil,
			wantType:   "*twelvedata.HTTPError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewHTTPError(tt.statusCode, tt.body, "https://api.twelvedata.com/test", tt.apiError, nil)

			// Check error type
			gotType := getErrorType(err)
			if gotType != tt.wantType {
				t.Errorf("NewHTTPError() type = %v, want %v", gotType, tt.wantType)
			}

			// Check that error implements error interface
			if err.Error() == "" {
				t.Errorf("NewHTTPError() error message is empty")
			}
		})
	}
}

func TestErrorTypeCheckers(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		checkers map[string]func(error) bool
		expected map[string]bool
	}{
		{
			name: "bad request error",
			err:  NewHTTPError(http.StatusBadRequest, []byte(""), "", nil, nil),
			checkers: map[string]func(error) bool{
				"IsBadRequestError":   IsBadRequestError,
				"IsUnauthorizedError": IsUnauthorizedError,
				"IsNotFoundError":     IsNotFoundError,
				"IsRateLimitError":    IsRateLimitError,
				"IsHTTPError":         IsHTTPError,
			},
			expected: map[string]bool{
				"IsBadRequestError":   true,
				"IsUnauthorizedError": false,
				"IsNotFoundError":     false,
				"IsRateLimitError":    false,
				"IsHTTPError":         true,
			},
		},
		{
			name: "unauthorized error",
			err:  NewHTTPError(http.StatusUnauthorized, []byte(""), "", nil, nil),
			checkers: map[string]func(error) bool{
				"IsBadRequestError":   IsBadRequestError,
				"IsUnauthorizedError": IsUnauthorizedError,
				"IsNotFoundError":     IsNotFoundError,
				"IsRateLimitError":    IsRateLimitError,
				"IsHTTPError":         IsHTTPError,
			},
			expected: map[string]bool{
				"IsBadRequestError":   false,
				"IsUnauthorizedError": true,
				"IsNotFoundError":     false,
				"IsRateLimitError":    false,
				"IsHTTPError":         true,
			},
		},
		{
			name: "timeout error",
			err:  &TimeoutError{Message: "Request timeout"},
			checkers: map[string]func(error) bool{
				"IsTimeoutError": IsTimeoutError,
				"IsNetworkError": IsNetworkError,
				"IsHTTPError":    IsHTTPError,
			},
			expected: map[string]bool{
				"IsTimeoutError": true,
				"IsNetworkError": false,
				"IsHTTPError":    false,
			},
		},
		{
			name: "network error",
			err:  &NetworkError{Message: "Connection refused", Cause: nil},
			checkers: map[string]func(error) bool{
				"IsTimeoutError": IsTimeoutError,
				"IsNetworkError": IsNetworkError,
				"IsHTTPError":    IsHTTPError,
			},
			expected: map[string]bool{
				"IsTimeoutError": false,
				"IsNetworkError": true,
				"IsHTTPError":    false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for checkerName, checker := range tt.checkers {
				result := checker(tt.err)
				expected := tt.expected[checkerName]

				if result != expected {
					t.Errorf("%s() = %v, want %v", checkerName, result, expected)
				}
			}
		})
	}
}

func TestErrorMessages(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		contains string
	}{
		{
			name: "bad request error with api error",
			err: &BadRequestError{
				HTTPError: HTTPError{StatusCode: 400, Message: "Bad Request"},
				APIError: &response.Error{
					Code:    400,
					Message: "Invalid parameter",
					Status:  "error",
				},
			},
			contains: "HTTP 400 Bad Request: code: 400, message: Invalid parameter, status: error",
		},
		{
			name: "unauthorized error without api error",
			err: &UnauthorizedError{
				HTTPError: HTTPError{StatusCode: 401, Message: "Unauthorized"},
				APIError:  nil,
			},
			contains: "HTTP 401 Unauthorized: Invalid API key or authentication failed",
		},
		{
			name: "timeout error",
			err: &TimeoutError{
				Message: "Request timeout after 30 seconds",
			},
			contains: "Request Timeout: Request timeout after 30 seconds",
		},
		{
			name: "network error with cause",
			err: &NetworkError{
				Message: "Connection failed",
				Cause:   http.ErrHandlerTimeout,
			},
			contains: "Network Error: Connection failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errMsg := tt.err.Error()
			if errMsg != tt.contains {
				t.Errorf("Error() = %q, want %q", errMsg, tt.contains)
			}
		})
	}
}

func TestNetworkErrorUnwrap(t *testing.T) {
	originalErr := http.ErrHandlerTimeout
	networkErr := &NetworkError{
		Message: "Connection failed",
		Cause:   originalErr,
	}

	if unwrapped := networkErr.Unwrap(); !errors.Is(unwrapped, originalErr) {
		t.Errorf("NetworkError.Unwrap() = %v, want %v", unwrapped, originalErr)
	}
}

func TestWebSocketErrors(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		checkers map[string]func(error) bool
		expected map[string]bool
	}{
		{
			name: "websocket connection error",
			err: &WSConnectionError{
				URL:     "wss://ws.twelvedata.com/v1/quotes/price",
				Message: "Failed to establish connection",
				Cause:   errors.New("connection refused"),
			},
			checkers: map[string]func(error) bool{
				"IsWSConnectionError": IsWSConnectionError,
				"IsWSMessageError":    IsWSMessageError,
				"IsWSError":           IsWSError,
				"IsHTTPError":         IsHTTPError,
			},
			expected: map[string]bool{
				"IsWSConnectionError": true,
				"IsWSMessageError":    false,
				"IsWSError":           true,
				"IsHTTPError":         false,
			},
		},
		{
			name: "websocket message error",
			err: &WSMessageError{
				Message: "Failed to write ping",
				Data:    []byte("ping data"),
				Cause:   errors.New("write timeout"),
			},
			checkers: map[string]func(error) bool{
				"IsWSConnectionError":   IsWSConnectionError,
				"IsWSMessageError":      IsWSMessageError,
				"IsWSSubscriptionError": IsWSSubscriptionError,
				"IsWSError":             IsWSError,
			},
			expected: map[string]bool{
				"IsWSConnectionError":   false,
				"IsWSMessageError":      true,
				"IsWSSubscriptionError": false,
				"IsWSError":             true,
			},
		},
		{
			name: "websocket subscription error",
			err: &WSSubscriptionError{
				Symbols: []string{"AAPL", "GOOGL"},
				Message: "Subscription failed",
				Cause:   errors.New("invalid symbols"),
			},
			checkers: map[string]func(error) bool{
				"IsWSSubscriptionError": IsWSSubscriptionError,
				"IsWSError":             IsWSError,
				"IsNetworkError":        IsNetworkError,
			},
			expected: map[string]bool{
				"IsWSSubscriptionError": true,
				"IsWSError":             true,
				"IsNetworkError":        false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for checkerName, checker := range tt.checkers {
				result := checker(tt.err)
				expected := tt.expected[checkerName]

				if result != expected {
					t.Errorf("%s() = %v, want %v for error: %v", checkerName, result, expected, tt.err)
				}
			}

			// Test error message format
			if tt.err.Error() == "" {
				t.Errorf("Error message should not be empty")
			}

			// Test Unwrap functionality if applicable
			if unwrapper, ok := tt.err.(interface{ Unwrap() error }); ok {
				if unwrapper.Unwrap() == nil {
					t.Errorf("Expected error to wrap another error, but Unwrap() returned nil")
				}
			}
		})
	}
}

func TestErrorContextAndURL(t *testing.T) {
	testURL := "https://api.twelvedata.com/stocks?apikey=test"

	tests := []struct {
		name        string
		err         error
		expectsURL  bool
		expectedURL string
	}{
		{
			name:        "http error with url context",
			err:         NewHTTPError(http.StatusBadRequest, []byte("test"), testURL, nil, nil),
			expectsURL:  true,
			expectedURL: testURL,
		},
		{
			name: "websocket connection error with url",
			err: &WSConnectionError{
				URL:     "wss://ws.twelvedata.com/v1/quotes/price",
				Message: "Connection failed",
				Cause:   errors.New("network error"),
			},
			expectsURL:  true,
			expectedURL: "wss://ws.twelvedata.com/v1/quotes/price",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errMsg := tt.err.Error()

			if tt.expectsURL && !strings.Contains(errMsg, tt.expectedURL) {
				t.Errorf("Expected error message to contain URL %q, got: %q", tt.expectedURL, errMsg)
			}
		})
	}
}

func TestErrorChaining(t *testing.T) {
	originalErr := errors.New("original network error")
	networkErr := &NetworkError{
		Message: "Connection failed",
		Cause:   originalErr,
	}
	wrappedErr := fmt.Errorf("service unavailable: %w", networkErr)

	// Test that we can find the original error through the chain
	if !errors.Is(wrappedErr, originalErr) {
		t.Errorf("Expected wrapped error to contain original error")
	}

	// Test that our error checkers work with wrapped errors
	if !IsNetworkError(wrappedErr) {
		t.Errorf("Expected IsNetworkError to work with wrapped errors")
	}

	// Test unwrapping chain
	var netErr *NetworkError
	if !errors.As(wrappedErr, &netErr) {
		t.Errorf("Expected errors.As to extract NetworkError from wrapped error")
	}

	if !errors.Is(netErr.Unwrap(), originalErr) {
		t.Errorf("Expected NetworkError.Unwrap() to return original error")
	}
}

func TestMalformedResponseHandling(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       []byte
		apiError   *response.Error
		url        string
	}{
		{
			name:       "malformed json body",
			statusCode: http.StatusBadRequest,
			body:       []byte("{invalid json"),
			apiError:   nil,
			url:        "https://api.twelvedata.com/test",
		},
		{
			name:       "empty body with error status",
			statusCode: http.StatusInternalServerError,
			body:       []byte(""),
			apiError:   nil,
			url:        "https://api.twelvedata.com/test",
		},
		{
			name:       "non-json error response",
			statusCode: http.StatusBadGateway,
			body:       []byte("Bad Gateway"),
			apiError:   nil,
			url:        "https://api.twelvedata.com/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewHTTPError(tt.statusCode, tt.body, tt.url, tt.apiError, nil)

			if err == nil {
				t.Errorf("Expected error, got nil")

				return
			}

			if !IsHTTPError(err) {
				t.Errorf("Expected IsHTTPError to return true")
			}

			errMsg := err.Error()
			if errMsg == "" {
				t.Errorf("Error message should not be empty")
			}

			// Verify URL is included in error message
			if !strings.Contains(errMsg, tt.url) {
				t.Errorf("Expected error message to contain URL %q, got: %q", tt.url, errMsg)
			}
		})
	}
}

func TestDomainErrorParsing(t *testing.T) {
	tests := []struct {
		name          string
		apiError      *response.Error
		statusCode    int
		url           string
		expectedType  string
		expectedCheck func(error) bool
		expectedMsg   string
	}{
		{
			name: "symbol not found error",
			apiError: &response.Error{
				Code:    400,
				Message: "**AAPL** not found: symbol may be delisted",
				Status:  "error",
			},
			statusCode:    400,
			url:           "https://api.twelvedata.com/stocks",
			expectedType:  "SymbolNotFoundError",
			expectedCheck: IsSymbolNotFoundError,
			expectedMsg:   "Symbol Not Found: AAPL - **AAPL** not found: symbol may be delisted",
		},
		{
			name: "new symbol not found error",
			apiError: &response.Error{
				Code:    400,
				Message: "**GOOGL** with specified criteria not found: check your parameters",
				Status:  "error",
			},
			statusCode:    400,
			url:           "https://api.twelvedata.com/stocks",
			expectedType:  "SymbolNotFoundError",
			expectedCheck: IsSymbolNotFoundError,
			expectedMsg:   "Symbol Not Found: GOOGL - **GOOGL** with specified criteria not found: check your parameters",
		},
		{
			name: "plan limitation error",
			apiError: &response.Error{
				Code:    403,
				Message: "Real-time data is not available with your plan",
				Status:  "error",
			},
			statusCode:    403,
			url:           "https://api.twelvedata.com/time_series",
			expectedType:  "PlanLimitationError",
			expectedCheck: IsPlanLimitationError,
			expectedMsg:   "Plan Limitation: Real-time data is not available with your plan",
		},
		{
			name: "insufficient credits error",
			apiError: &response.Error{
				Code:    402,
				Message: "insufficient credits to complete this request",
				Status:  "error",
			},
			statusCode:    402,
			url:           "https://api.twelvedata.com/time_series",
			expectedType:  "InsufficientCreditsError",
			expectedCheck: IsInsufficientCreditsError,
			expectedMsg:   "Insufficient Credits: insufficient credits to complete this request",
		},
		{
			name: "invalid api key error",
			apiError: &response.Error{
				Code:    401,
				Message: "invalid api key provided",
				Status:  "error",
			},
			statusCode:    401,
			url:           "https://api.twelvedata.com/stocks",
			expectedType:  "APIKeyError",
			expectedCheck: IsAPIKeyError,
			expectedMsg:   "API Key Error: Invalid API key provided",
		},
		{
			name: "api key required error",
			apiError: &response.Error{
				Code:    401,
				Message: "api key is required for this endpoint",
				Status:  "error",
			},
			statusCode:    401,
			url:           "https://api.twelvedata.com/stocks",
			expectedType:  "APIKeyError",
			expectedCheck: IsAPIKeyError,
			expectedMsg:   "API Key Error: API key is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ParseDomainError(tt.apiError, tt.statusCode, tt.url)

			if err == nil {
				t.Errorf("Expected domain error, got nil")

				return
			}

			// Test error type checking
			if !tt.expectedCheck(err) {
				t.Errorf("Expected %s error checker to return true, got false", tt.expectedType)
			}

			// Test error message
			if err.Error() != tt.expectedMsg {
				t.Errorf("Expected error message %q, got %q", tt.expectedMsg, err.Error())
			}

			// Test that it's recognized as a domain error
			if !IsDomainError(err) {
				t.Errorf("Expected IsDomainError to return true")
			}

			// Test error unwrapping
			if unwrapper, ok := err.(interface{ Unwrap() error }); ok {
				if unwrapped := unwrapper.Unwrap(); unwrapped == nil {
					t.Errorf("Expected error to wrap another error")
				}
			}
		})
	}
}

func TestDomainErrorCheckers(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		checkers map[string]func(error) bool
		expected map[string]bool
	}{
		{
			name: "symbol not found error",
			err: &SymbolNotFoundError{
				Symbol:  "AAPL",
				Message: "Symbol not found",
			},
			checkers: map[string]func(error) bool{
				"IsSymbolNotFoundError":      IsSymbolNotFoundError,
				"IsPlanLimitationError":      IsPlanLimitationError,
				"IsInsufficientCreditsError": IsInsufficientCreditsError,
				"IsAPIKeyError":              IsAPIKeyError,
				"IsDomainError":              IsDomainError,
			},
			expected: map[string]bool{
				"IsSymbolNotFoundError":      true,
				"IsPlanLimitationError":      false,
				"IsInsufficientCreditsError": false,
				"IsAPIKeyError":              false,
				"IsDomainError":              true,
			},
		},
		{
			name: "plan limitation error",
			err: &PlanLimitationError{
				Feature: "Real-time data",
				Plan:    "basic",
				Message: "Not available",
			},
			checkers: map[string]func(error) bool{
				"IsSymbolNotFoundError":      IsSymbolNotFoundError,
				"IsPlanLimitationError":      IsPlanLimitationError,
				"IsInsufficientCreditsError": IsInsufficientCreditsError,
				"IsDomainError":              IsDomainError,
			},
			expected: map[string]bool{
				"IsSymbolNotFoundError":      false,
				"IsPlanLimitationError":      true,
				"IsInsufficientCreditsError": false,
				"IsDomainError":              true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for checkerName, checker := range tt.checkers {
				result := checker(tt.err)
				expected := tt.expected[checkerName]

				if result != expected {
					t.Errorf("%s() = %v, want %v", checkerName, result, expected)
				}
			}
		})
	}
}

func TestErrorMessageExtraction(t *testing.T) {
	tests := []struct {
		name      string
		message   string
		extractor func(string) string
		expected  string
	}{
		{
			name:      "extract symbol from standard message",
			message:   "**AAPL** not found: symbol may be delisted",
			extractor: extractSymbolFromMessage,
			expected:  "AAPL",
		},
		{
			name:      "extract symbol from criteria message",
			message:   "**GOOGL** with specified criteria not found",
			extractor: extractSymbolFromMessage,
			expected:  "GOOGL",
		},
		{
			name:      "extract symbol - no symbol found",
			message:   "Invalid request parameters",
			extractor: extractSymbolFromMessage,
			expected:  "",
		},
		{
			name:      "extract feature from plan message",
			message:   "Real-time data is not available with your plan",
			extractor: extractFeatureFromMessage,
			expected:  "Real-time data",
		},
		{
			name:      "extract feature with prefix",
			message:   "The advanced analytics is not available with your plan",
			extractor: extractFeatureFromMessage,
			expected:  "advanced analytics",
		},
		{
			name:      "extract feature - no feature found",
			message:   "Some other error message",
			extractor: extractFeatureFromMessage,
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.extractor(tt.message)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestWrappedDomainErrors(t *testing.T) {
	originalSymbolErr := &SymbolNotFoundError{
		Symbol:  "TSLA",
		Message: "Symbol not found",
	}
	wrappedErr := fmt.Errorf("request failed: %w", originalSymbolErr)

	// Test that domain error checkers work with wrapped errors
	if !IsSymbolNotFoundError(wrappedErr) {
		t.Errorf("Expected IsSymbolNotFoundError to work with wrapped errors")
	}

	if !IsDomainError(wrappedErr) {
		t.Errorf("Expected IsDomainError to work with wrapped errors")
	}

	// Test that we can extract the original error
	var symbolErr *SymbolNotFoundError
	if !errors.As(wrappedErr, &symbolErr) {
		t.Errorf("Expected errors.As to extract SymbolNotFoundError")
	}

	if symbolErr.Symbol != "TSLA" {
		t.Errorf("Expected symbol TSLA, got %s", symbolErr.Symbol)
	}
}

func TestWrappedErrorHandling(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		checkers map[string]func(error) bool
		expected map[string]bool
	}{
		{
			name: "wrapped bad request error",
			err:  fmt.Errorf("wrapper: %w", NewHTTPError(http.StatusBadRequest, []byte(""), "", nil, nil)),
			checkers: map[string]func(error) bool{
				"IsBadRequestError": IsBadRequestError,
				"IsHTTPError":       IsHTTPError,
				"IsTimeoutError":    IsTimeoutError,
			},
			expected: map[string]bool{
				"IsBadRequestError": true,
				"IsHTTPError":       true,
				"IsTimeoutError":    false,
			},
		},
		{
			name: "wrapped timeout error",
			err:  fmt.Errorf("connection failed: %w", &TimeoutError{Message: "Request timeout"}),
			checkers: map[string]func(error) bool{
				"IsTimeoutError": IsTimeoutError,
				"IsNetworkError": IsNetworkError,
				"IsHTTPError":    IsHTTPError,
			},
			expected: map[string]bool{
				"IsTimeoutError": true,
				"IsNetworkError": false,
				"IsHTTPError":    false,
			},
		},
		{
			name: "wrapped network error",
			err:  fmt.Errorf("service unavailable: %w", &NetworkError{Message: "Connection refused", Cause: nil}),
			checkers: map[string]func(error) bool{
				"IsNetworkError": IsNetworkError,
				"IsTimeoutError": IsTimeoutError,
				"IsHTTPError":    IsHTTPError,
			},
			expected: map[string]bool{
				"IsNetworkError": true,
				"IsTimeoutError": false,
				"IsHTTPError":    false,
			},
		},
		{
			name: "double wrapped error",
			err:  fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", NewHTTPError(http.StatusUnauthorized, []byte(""), "", nil, nil))),
			checkers: map[string]func(error) bool{
				"IsUnauthorizedError": IsUnauthorizedError,
				"IsHTTPError":         IsHTTPError,
				"IsBadRequestError":   IsBadRequestError,
			},
			expected: map[string]bool{
				"IsUnauthorizedError": true,
				"IsHTTPError":         true,
				"IsBadRequestError":   false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for checkerName, checker := range tt.checkers {
				result := checker(tt.err)
				expected := tt.expected[checkerName]

				if result != expected {
					t.Errorf("%s() = %v, want %v for wrapped error: %v", checkerName, result, expected, tt.err)
				}
			}
		})
	}
}

// Helper function to get error type name for testing.
func getErrorType(err error) string {
	switch err.(type) {
	case *BadRequestError:
		return "*twelvedata.BadRequestError"
	case *UnauthorizedError:
		return "*twelvedata.UnauthorizedError"
	case *NotFoundError:
		return "*twelvedata.NotFoundError"
	case *TooManyRequestsError:
		return "*twelvedata.TooManyRequestsError"
	case *InternalServerError:
		return "*twelvedata.InternalServerError"
	case *HTTPError:
		return "*twelvedata.HTTPError"
	case *TimeoutError:
		return "*twelvedata.TimeoutError"
	case *NetworkError:
		return "*twelvedata.NetworkError"
	default:
		return "unknown"
	}
}
