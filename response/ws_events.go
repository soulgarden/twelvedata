package response

import "github.com/guregu/null/v6"

// WSEventType represents the type of WebSocket event.
type WSEventType string

const (
	// WSEventPrice represents a real-time price update event.
	WSEventPrice WSEventType = "price"
	// WSEventSubscribeStatus represents a subscription status event.
	WSEventSubscribeStatus WSEventType = "subscribe-status"
	// WSEventError represents an error event.
	WSEventError WSEventType = "error"
)

// WSBaseEvent represents the base structure for all WebSocket events.
type WSBaseEvent struct {
	Event WSEventType `json:"event"`
}

// WSPriceEvent represents a real-time price event with comprehensive field support.
// Supports stocks, forex, cryptocurrencies, and other instrument types.
type WSPriceEvent struct {
	Event     WSEventType `json:"event"`
	Symbol    string      `json:"symbol"`
	Currency  string      `json:"currency"`
	Exchange  string      `json:"exchange"`
	MicCode   string      `json:"mic_code,omitempty"`
	Type      string      `json:"type"`
	Timestamp int64       `json:"timestamp"`
	Price     float64     `json:"price"`

	// Stock-specific fields
	DayVolume null.Int `json:"day_volume,omitempty"`

	// Forex/Crypto-specific fields
	Bid           null.Float  `json:"bid,omitempty"`
	Ask           null.Float  `json:"ask,omitempty"`
	CurrencyBase  null.String `json:"currency_base,omitempty"`
	CurrencyQuote null.String `json:"currency_quote,omitempty"`
}

// WSSubscribeStatusEvent represents a subscription status response.
type WSSubscribeStatusEvent struct {
	Event   WSEventType            `json:"event"`
	Status  string                 `json:"status"`
	Success []WSSubscriptionResult `json:"success,omitempty"`
	Fails   []WSSubscriptionFail   `json:"fails,omitempty"`
}

// WSSubscriptionResult represents a successful subscription result.
type WSSubscriptionResult struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	Country  string `json:"country"`
	Type     string `json:"type"`
}

// WSSubscriptionFail represents a failed subscription attempt.
type WSSubscriptionFail struct {
	Symbol  string `json:"symbol"`
	Message string `json:"message"`
}

// WSErrorEvent represents an error event from the WebSocket server.
type WSErrorEvent struct {
	Event   WSEventType `json:"event"`
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message"`
	Status  string      `json:"status,omitempty"`
}

// WSEvent is a union type that can represent any WebSocket event.
// Use type assertion or switch on Event field to determine the specific type.
type WSEvent struct {
	Event WSEventType `json:"event"`
	// Embed anonymous structs to allow for flexible unmarshaling
	*WSPriceEvent
	*WSSubscribeStatusEvent
	*WSErrorEvent
}
