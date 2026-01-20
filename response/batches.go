package response

import "github.com/guregu/null/v6"

const (
	// StatusSuccess represents the successful status value returned in batch API responses.
	StatusSuccess = "success"
)

// Batches represents the response from the batch API endpoint.
type Batches struct {
	Code   int                      `json:"code"`
	Status string                   `json:"status"`
	Data   map[string]BatchResponse `json:"data"`
}

// BatchResponse represents a single response within a batch operation.
type BatchResponse struct {
	Status   string            `json:"status"`
	Response BatchResponseData `json:"response"`
}

// BatchResponseData represents the actual data returned for each batch request.
// This can contain either successful response data or error information.
type BatchResponseData struct {
	// Success response fields - these will be populated for successful requests
	Meta   *BatchMeta   `json:"meta,omitempty"`
	Values []BatchValue `json:"values,omitempty"`
	Status null.String  `json:"status,omitempty"`

	// Quote-specific fields for /quote endpoint responses
	Symbol                null.String `json:"symbol,omitempty"`
	Name                  null.String `json:"name,omitempty"`
	Exchange              null.String `json:"exchange,omitempty"`
	MicCode               null.String `json:"mic_code,omitempty"`
	Currency              null.String `json:"currency,omitempty"`
	Datetime              null.String `json:"datetime,omitempty"`
	Timestamp             null.Int    `json:"timestamp,omitempty"`
	LastQuoteAt           null.Int    `json:"last_quote_at,omitempty"`
	Open                  null.String `json:"open,omitempty"`
	High                  null.String `json:"high,omitempty"`
	Low                   null.String `json:"low,omitempty"`
	Close                 null.String `json:"close,omitempty"`
	Volume                null.String `json:"volume,omitempty"`
	PreviousClose         null.String `json:"previous_close,omitempty"`
	Change                null.String `json:"change,omitempty"`
	PercentChange         null.String `json:"percent_change,omitempty"`
	AverageVolume         null.String `json:"average_volume,omitempty"`
	Rolling1DChange       null.String `json:"rolling_1d_change,omitempty"`
	Rolling7DChange       null.String `json:"rolling_7d_change,omitempty"`
	RollingChange         null.String `json:"rolling_change,omitempty"`
	IsMarketOpen          null.Bool   `json:"is_market_open,omitempty"`
	ExtendedChange        null.String `json:"extended_change,omitempty"`
	ExtendedPercentChange null.String `json:"extended_percent_change,omitempty"`
	ExtendedPrice         null.String `json:"extended_price,omitempty"`
	ExtendedTimestamp     null.String `json:"extended_timestamp,omitempty"`

	// Error response fields - these will be populated for failed requests
	Code    null.Int    `json:"code,omitempty"`
	Message null.String `json:"message,omitempty"`

	// Exchange rate specific fields
	Rate     null.Float  `json:"rate,omitempty"`
	FromCode null.String `json:"from_code,omitempty"`
	ToCode   null.String `json:"to_code,omitempty"`
	Amount   null.Float  `json:"amount,omitempty"`
}

// BatchMeta represents metadata for batch responses, typically from time series or similar endpoints.
type BatchMeta struct {
	Symbol           null.String `json:"symbol,omitempty"`
	Interval         null.String `json:"interval,omitempty"`
	Currency         null.String `json:"currency,omitempty"`
	ExchangeTimezone null.String `json:"exchange_timezone,omitempty"`
	Exchange         null.String `json:"exchange,omitempty"`
	MicCode          null.String `json:"mic_code,omitempty"`
	Type             null.String `json:"type,omitempty"`
}

// BatchValue represents time series values in batch responses.
type BatchValue struct {
	Datetime null.String `json:"datetime,omitempty"`
	Open     null.String `json:"open,omitempty"`
	High     null.String `json:"high,omitempty"`
	Low      null.String `json:"low,omitempty"`
	Close    null.String `json:"close,omitempty"`
	Volume   null.String `json:"volume,omitempty"`
}

// GetRequestResponse returns the response data for a specific request ID.
func (b *Batches) GetRequestResponse(requestID string) (BatchResponse, bool) {
	response, exists := b.Data[requestID]
	return response, exists
}

// IsSuccess returns true if the batch request was successful.
func (b *Batches) IsSuccess() bool {
	return b.Code == 200 && b.Status == StatusSuccess
}

// HasErrors returns true if any of the individual requests in the batch failed.
func (b *Batches) HasErrors() bool {
	for _, response := range b.Data {
		if response.Status != StatusSuccess || response.Response.Code.Valid {
			return true
		}
	}
	return false
}

// GetErrors returns a map of request IDs to their error messages for failed requests.
func (b *Batches) GetErrors() map[string]string {
	errors := make(map[string]string)
	for requestID, response := range b.Data {
		if response.Status != "success" && response.Response.Message.Valid {
			errors[requestID] = response.Response.Message.String
		} else if response.Response.Code.Valid && response.Response.Code.Int64 != 200 {
			if response.Response.Message.Valid {
				errors[requestID] = response.Response.Message.String
			} else {
				errors[requestID] = "Unknown error occurred"
			}
		}
	}
	return errors
}
