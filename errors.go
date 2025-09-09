package twelvedata

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/soulgarden/twelvedata/dictionary"
	"github.com/soulgarden/twelvedata/response"
)

// HTTPError represents HTTP-related errors with status codes.
type HTTPError struct {
	StatusCode int
	Message    string
	Body       []byte
	URL        string // request URL that failed
	Cause      error  // underlying error that caused this HTTP error
}

func (e HTTPError) Error() string {
	if e.URL != "" {
		return fmt.Sprintf("HTTP %d %s (URL: %s)", e.StatusCode, e.Message, e.URL)
	}

	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

// Unwrap returns the underlying cause if this HTTPError wraps another error.
func (e HTTPError) Unwrap() error {
	return e.Cause
}

// IsHTTPError checks if an error is an HTTPError or any of its subtypes.
func IsHTTPError(err error) bool {
	// Check for direct HTTPError
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		return true
	}

	// Check for specific HTTP error types
	var badReqErr *BadRequestError

	var unauthorizedErr *UnauthorizedError

	var notFoundErr *NotFoundError

	var rateLimitErr *TooManyRequestsError

	var serverErr *InternalServerError

	return errors.As(err, &badReqErr) ||
		errors.As(err, &unauthorizedErr) ||
		errors.As(err, &notFoundErr) ||
		errors.As(err, &rateLimitErr) ||
		errors.As(err, &serverErr)
}

// BadRequestError represents 400 Bad Request errors.
type BadRequestError struct {
	HTTPError
	APIError *response.Error
}

func (e BadRequestError) Error() string {
	baseMsg := "HTTP 400 Bad Request"
	if e.APIError != nil {
		baseMsg += ": " + e.APIError.Error()
	} else {
		baseMsg += ": " + e.Message
	}

	if e.URL != "" {
		baseMsg += " (URL: " + e.URL + ")"
	}

	return baseMsg
}

// UnauthorizedError represents 401 Unauthorized errors.
type UnauthorizedError struct {
	HTTPError
	APIError *response.Error
}

func (e UnauthorizedError) Error() string {
	baseMsg := "HTTP 401 Unauthorized"
	if e.APIError != nil {
		baseMsg += ": " + e.APIError.Error()
	} else {
		baseMsg += ": Invalid API key or authentication failed"
	}

	if e.URL != "" {
		baseMsg += " (URL: " + e.URL + ")"
	}

	return baseMsg
}

// NotFoundError represents 404 Not Found errors.
type NotFoundError struct {
	HTTPError
	APIError *response.Error
}

func (e NotFoundError) Error() string {
	baseMsg := "HTTP 404 Not Found"
	if e.APIError != nil {
		baseMsg += ": " + e.APIError.Error()
	} else {
		baseMsg += ": " + e.Message
	}

	if e.URL != "" {
		baseMsg += " (URL: " + e.URL + ")"
	}

	return baseMsg
}

// TooManyRequestsError represents 429 Rate Limit errors.
type TooManyRequestsError struct {
	HTTPError
	APIError *response.Error
}

func (e TooManyRequestsError) Error() string {
	baseMsg := "HTTP 429 Too Many Requests"
	if e.APIError != nil {
		baseMsg += ": " + e.APIError.Error()
	} else {
		baseMsg += ": Rate limit exceeded"
	}

	if e.URL != "" {
		baseMsg += " (URL: " + e.URL + ")"
	}

	return baseMsg
}

// InternalServerError represents 500 Internal Server errors.
type InternalServerError struct {
	HTTPError
	APIError *response.Error
}

func (e InternalServerError) Error() string {
	baseMsg := "HTTP 500 Internal Server Error"
	if e.APIError != nil {
		baseMsg += ": " + e.APIError.Error()
	} else {
		baseMsg += ": " + e.Message
	}

	if e.URL != "" {
		baseMsg += " (URL: " + e.URL + ")"
	}

	return baseMsg
}

// TimeoutError represents request timeout errors.
type TimeoutError struct {
	Message string
}

func (e TimeoutError) Error() string {
	return "Request Timeout: " + e.Message
}

// NetworkError represents network connectivity errors.
type NetworkError struct {
	Message string
	Cause   error
}

func (e NetworkError) Error() string {
	return "Network Error: " + e.Message
}

func (e NetworkError) Unwrap() error {
	return e.Cause
}

// WebSocket-specific errors

// WSConnectionError represents WebSocket connection errors.
type WSConnectionError struct {
	URL     string
	Message string
	Cause   error
}

func (e WSConnectionError) Error() string {
	return fmt.Sprintf("WebSocket Connection Error: %s (URL: %s)", e.Message, e.URL)
}

func (e WSConnectionError) Unwrap() error {
	return e.Cause
}

// WSMessageError represents WebSocket message handling errors.
type WSMessageError struct {
	Message string
	Data    []byte
	Cause   error
}

func (e WSMessageError) Error() string {
	return "WebSocket Message Error: " + e.Message
}

func (e WSMessageError) Unwrap() error {
	return e.Cause
}

// WSSubscriptionError represents WebSocket subscription errors.
type WSSubscriptionError struct {
	Symbols []string
	Message string
	Cause   error
}

func (e WSSubscriptionError) Error() string {
	return fmt.Sprintf("WebSocket Subscription Error: %s (symbols: %v)", e.Message, e.Symbols)
}

func (e WSSubscriptionError) Unwrap() error {
	return e.Cause
}

// NewHTTPError creates appropriate typed error based on HTTP status code.
func NewHTTPError(statusCode int, body []byte, url string, apiError *response.Error, cause error) error {
	baseError := HTTPError{
		StatusCode: statusCode,
		Body:       body,
		URL:        url,
		Message:    http.StatusText(statusCode),
		Cause:      cause,
	}

	switch statusCode {
	case http.StatusBadRequest:
		return &BadRequestError{
			HTTPError: baseError,
			APIError:  apiError,
		}
	case http.StatusUnauthorized:
		return &UnauthorizedError{
			HTTPError: baseError,
			APIError:  apiError,
		}
	case http.StatusNotFound:
		return &NotFoundError{
			HTTPError: baseError,
			APIError:  apiError,
		}
	case http.StatusTooManyRequests:
		return &TooManyRequestsError{
			HTTPError: baseError,
			APIError:  apiError,
		}
	case http.StatusInternalServerError:
		return &InternalServerError{
			HTTPError: baseError,
			APIError:  apiError,
		}
	default:
		if apiError != nil {
			baseError.Message = apiError.Error()
		}

		return &baseError
	}
}

// Error type checking helpers.

// IsBadRequestError checks if an error is a BadRequestError type.
func IsBadRequestError(err error) bool {
	var badReqErr *BadRequestError

	return errors.As(err, &badReqErr)
}

// IsUnauthorizedError checks if an error is an UnauthorizedError type.
func IsUnauthorizedError(err error) bool {
	var unauthorizedErr *UnauthorizedError

	return errors.As(err, &unauthorizedErr)
}

// IsNotFoundError checks if an error is a NotFoundError type.
func IsNotFoundError(err error) bool {
	var notFoundErr *NotFoundError

	return errors.As(err, &notFoundErr)
}

// IsRateLimitError checks if an error is a TooManyRequestsError type.
func IsRateLimitError(err error) bool {
	var rateLimitErr *TooManyRequestsError

	return errors.As(err, &rateLimitErr)
}

// IsTimeoutError checks if an error is a TimeoutError type.
func IsTimeoutError(err error) bool {
	var timeoutErr *TimeoutError

	return errors.As(err, &timeoutErr)
}

// IsNetworkError checks if an error is a NetworkError type.
func IsNetworkError(err error) bool {
	var networkErr *NetworkError

	return errors.As(err, &networkErr)
}

// WebSocket error checking functions

// IsWSConnectionError checks if an error is a WSConnectionError type.
func IsWSConnectionError(err error) bool {
	var wsConnErr *WSConnectionError

	return errors.As(err, &wsConnErr)
}

// IsWSMessageError checks if an error is a WSMessageError type.
func IsWSMessageError(err error) bool {
	var wsMsgErr *WSMessageError

	return errors.As(err, &wsMsgErr)
}

// IsWSSubscriptionError checks if an error is a WSSubscriptionError type.
func IsWSSubscriptionError(err error) bool {
	var wsSubErr *WSSubscriptionError

	return errors.As(err, &wsSubErr)
}

// IsWSError checks if an error is any WebSocket-related error type.
func IsWSError(err error) bool {
	return IsWSConnectionError(err) || IsWSMessageError(err) || IsWSSubscriptionError(err)
}

// Domain-specific errors for Twelve Data API

// SymbolNotFoundError represents errors when a requested symbol is not found.
type SymbolNotFoundError struct {
	Symbol  string
	Message string
	Cause   error
}

func (e SymbolNotFoundError) Error() string {
	if e.Symbol != "" {
		return fmt.Sprintf("Symbol Not Found: %s - %s", e.Symbol, e.Message)
	}

	return "Symbol Not Found: " + e.Message
}

func (e SymbolNotFoundError) Unwrap() error {
	return e.Cause
}

// PlanLimitationError represents errors when a feature is not available with the current plan.
type PlanLimitationError struct {
	Feature string
	Plan    string
	Message string
	Cause   error
}

func (e PlanLimitationError) Error() string {
	if e.Feature != "" && e.Plan != "" {
		return fmt.Sprintf("Plan Limitation: %s is not available with %s plan", e.Feature, e.Plan)
	}

	return "Plan Limitation: " + e.Message
}

func (e PlanLimitationError) Unwrap() error {
	return e.Cause
}

// InsufficientCreditsError represents errors when user has insufficient API credits.
type InsufficientCreditsError struct {
	Required  int64
	Available int64
	Message   string
	Cause     error
}

func (e InsufficientCreditsError) Error() string {
	if e.Required > 0 && e.Available >= 0 {
		return fmt.Sprintf("Insufficient Credits: required %d, available %d", e.Required, e.Available)
	}

	return "Insufficient Credits: " + e.Message
}

func (e InsufficientCreditsError) Unwrap() error {
	return e.Cause
}

// APIKeyError represents API key related errors.
type APIKeyError struct {
	Type    string // "invalid", "required", "expired"
	Message string
	Cause   error
}

func (e APIKeyError) Error() string {
	switch e.Type {
	case "invalid":
		return "API Key Error: Invalid API key provided"
	case "required":
		return "API Key Error: API key is required"
	case "expired":
		return "API Key Error: API key has expired"
	default:
		return "API Key Error: " + e.Message
	}
}

func (e APIKeyError) Unwrap() error {
	return e.Cause
}

// Domain-specific error checking functions

// IsSymbolNotFoundError checks if an error is a SymbolNotFoundError type.
func IsSymbolNotFoundError(err error) bool {
	var symbolErr *SymbolNotFoundError

	return errors.As(err, &symbolErr)
}

// IsPlanLimitationError checks if an error is a PlanLimitationError type.
func IsPlanLimitationError(err error) bool {
	var planErr *PlanLimitationError

	return errors.As(err, &planErr)
}

// IsInsufficientCreditsError checks if an error is an InsufficientCreditsError type.
func IsInsufficientCreditsError(err error) bool {
	var creditsErr *InsufficientCreditsError

	return errors.As(err, &creditsErr)
}

// IsAPIKeyError checks if an error is an APIKeyError type.
func IsAPIKeyError(err error) bool {
	var keyErr *APIKeyError

	return errors.As(err, &keyErr)
}

// IsDomainError checks if an error is any of the Twelve Data domain-specific errors.
func IsDomainError(err error) bool {
	return IsSymbolNotFoundError(err) ||
		IsPlanLimitationError(err) ||
		IsInsufficientCreditsError(err) ||
		IsAPIKeyError(err)
}

// API error message parsing functions

// ParseDomainError analyzes an API error response and converts it to a domain-specific error.
func ParseDomainError(apiError *response.Error, _ int, _ string) error {
	if apiError == nil || apiError.Message == "" {
		return nil
	}

	message := strings.ToLower(apiError.Message)

	// Check for symbol not found errors
	if strings.Contains(message, "not found:") && strings.Contains(message, "**") ||
		strings.Contains(message, "with specified criteria not found") && strings.Contains(message, "**") {
		symbol := extractSymbolFromMessage(apiError.Message)

		return &SymbolNotFoundError{
			Symbol:  symbol,
			Message: apiError.Message,
			Cause:   fmt.Errorf("API returned symbol not found: %s", apiError.Message),
		}
	}

	// Check for plan limitation errors (old format)
	if strings.Contains(message, strings.ToLower(dictionary.IsNotAvailableWithYourPlanMsg)) {
		feature := extractFeatureFromMessage(apiError.Message)

		return &PlanLimitationError{
			Feature: feature,
			Message: apiError.Message,
			Cause:   fmt.Errorf("API returned plan limitation: %s", apiError.Message),
		}
	}

	// Check for plan limitation errors (new format: "is available exclusively with")
	if strings.Contains(message, strings.ToLower(dictionary.IsAvailableExclusivelyMsg)) {
		feature := extractFeatureFromExclusiveMessage(apiError.Message)

		return &PlanLimitationError{
			Feature: feature,
			Message: apiError.Message,
			Cause:   fmt.Errorf("API returned plan limitation: %s", apiError.Message),
		}
	}

	// Check for demo API key limitation errors
	if strings.Contains(message, strings.ToLower(dictionary.DemoAPIKeyLimitationMsg)) {
		return &PlanLimitationError{
			Feature: "demo API key",
			Message: apiError.Message,
			Cause:   fmt.Errorf("API returned demo key limitation: %s", apiError.Message),
		}
	}

	// Check for insufficient credits errors
	if strings.Contains(message, strings.ToLower(dictionary.InsufficientCreditsMsg)) {
		return &InsufficientCreditsError{
			Message: apiError.Message,
			Cause:   fmt.Errorf("API returned insufficient credits: %s", apiError.Message),
		}
	}

	// Check for daily credit limit exhaustion (new pattern)
	if strings.Contains(message, "you have run out of api credits") {
		return &InsufficientCreditsError{
			Message: apiError.Message,
			Cause:   fmt.Errorf("API returned daily credit limit exhausted: %s", apiError.Message),
		}
	}

	// Check for API key errors
	if strings.Contains(message, strings.ToLower(dictionary.APIKeyInvalidMsg)) {
		return &APIKeyError{
			Type:    "invalid",
			Message: apiError.Message,
			Cause:   fmt.Errorf("API returned invalid key: %s", apiError.Message),
		}
	}

	if strings.Contains(message, strings.ToLower(dictionary.APIKeyRequiredMsg)) {
		return &APIKeyError{
			Type:    "required",
			Message: apiError.Message,
			Cause:   fmt.Errorf("API returned key required: %s", apiError.Message),
		}
	}

	return nil
}

// extractSymbolFromMessage tries to extract the symbol name from error messages like "**AAPL** not found:".
func extractSymbolFromMessage(message string) string {
	// Look for pattern **SYMBOL**
	start := strings.Index(message, "**")
	if start == -1 {
		return ""
	}

	start += 2

	end := strings.Index(message[start:], "**")
	if end == -1 {
		return ""
	}

	return message[start : start+end]
}

// extractFeatureFromMessage tries to extract the feature name from plan limitation messages.
func extractFeatureFromMessage(message string) string {
	// Look for text before " is not available with your plan"
	planMsg := strings.ToLower(dictionary.IsNotAvailableWithYourPlanMsg)

	idx := strings.LastIndex(strings.ToLower(message), planMsg)
	if idx == -1 {
		return ""
	}

	feature := strings.TrimSpace(message[:idx])
	// Remove common prefixes
	feature = strings.TrimPrefix(feature, "The ")
	feature = strings.TrimPrefix(feature, "This ")
	feature = strings.TrimPrefix(feature, "Feature ")

	return feature
}

// extractFeatureFromExclusiveMessage extracts feature name from new format messages like
// "/dividends_calendar is available exclusively with grow or pro or ultra or enterprise plans.".
func extractFeatureFromExclusiveMessage(message string) string {
	exclusiveMsg := strings.ToLower(dictionary.IsAvailableExclusivelyMsg)
	idx := strings.Index(strings.ToLower(message), exclusiveMsg)
	if idx == -1 {
		return ""
	}

	feature := strings.TrimSpace(message[:idx])
	// Remove leading slash if present
	feature = strings.TrimPrefix(feature, "/")
	// Replace underscores with spaces for readability
	feature = strings.ReplaceAll(feature, "_", " ")

	return feature
}

// WrapWithDomainError wraps an existing error with domain-specific error if applicable.
func WrapWithDomainError(err error, apiError *response.Error, statusCode int, url string) error {
	if domainErr := ParseDomainError(apiError, statusCode, url); domainErr != nil {
		return domainErr
	}

	return err
}
