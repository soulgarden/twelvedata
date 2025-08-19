package request

// WSAction represents the action types for WebSocket messages.
type WSAction string

const (
	// WSActionSubscribe subscribes to price updates for specified symbols.
	WSActionSubscribe WSAction = "subscribe"
	// WSActionUnsubscribe unsubscribes from price updates for specified symbols.
	WSActionUnsubscribe WSAction = "unsubscribe"
	// WSActionReset resets the connection by unsubscribing from all symbols.
	WSActionReset WSAction = "reset"
	// WSActionHeartbeat sends a heartbeat to maintain connection stability.
	WSActionHeartbeat WSAction = "heartbeat"
)

// WSSubscribeRequest represents a WebSocket subscription request with symbols.
type WSSubscribeRequest struct {
	Action WSAction           `json:"action"`
	Params *WSSubscribeParams `json:"params,omitempty"`
}

// WSSubscribeParams contains the parameters for subscription requests.
type WSSubscribeParams struct {
	Symbols interface{} `json:"symbols"`
}

// WSSymbolExtended represents an extended format symbol with additional parameters.
type WSSymbolExtended struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange,omitempty"`
	MicCode  string `json:"mic_code,omitempty"`
	Type     string `json:"type,omitempty"`
}

// WSHeartbeatRequest represents a heartbeat message to maintain connection stability.
type WSHeartbeatRequest struct {
	Action WSAction `json:"action"`
}

// WSResetRequest represents a reset message to clear all subscriptions.
type WSResetRequest struct {
	Action WSAction `json:"action"`
}
