package twelvedata

import (
	"encoding/json"
	"fmt"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/response"
)

// wsMessageParser provides optimized JSON parsing for WebSocket messages using union types.
// This approach parses the entire message once and determines the event type, avoiding
// the need for separate unmarshaling operations.
type wsMessageParser struct {
	union wsEventUnion
}

// wsEventUnion represents a union of all possible WebSocket events.
// Only one field will be populated after parsing.
type wsEventUnion struct {
	Event response.WSEventType `json:"event"`

	// Price event fields (embedded inline for efficient parsing)
	Symbol        string      `json:"symbol,omitempty"`
	Currency      string      `json:"currency,omitempty"`
	Exchange      string      `json:"exchange,omitempty"`
	MicCode       string      `json:"mic_code,omitempty"`
	Type          string      `json:"type,omitempty"`
	Timestamp     null.Int    `json:"timestamp,omitempty"`
	Price         null.Float  `json:"price,omitempty"`
	DayVolume     null.Int    `json:"day_volume,omitempty"`
	Bid           null.Float  `json:"bid,omitempty"`
	Ask           null.Float  `json:"ask,omitempty"`
	CurrencyBase  null.String `json:"currency_base,omitempty"`
	CurrencyQuote null.String `json:"currency_quote,omitempty"`

	// Status event fields
	Status  string                          `json:"status,omitempty"`
	Success []response.WSSubscriptionResult `json:"success,omitempty"`
	Fails   []response.WSSubscriptionFail   `json:"fails,omitempty"`

	// Error event fields
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// newWSMessageParser creates a new parser instance.
func newWSMessageParser() *wsMessageParser {
	return &wsMessageParser{}
}

// parseMessage performs single-pass JSON parsing to extract both event type and data.
func (p *wsMessageParser) parseMessage(message []byte) (response.WSEventType, error) {
	p.union = wsEventUnion{}
	if err := json.Unmarshal(message, &p.union); err != nil {
		return "", fmt.Errorf("failed to parse message: %w", err)
	}
	return p.union.Event, nil
}

// getPriceEvent converts the union to a WSPriceEvent.
func (p *wsMessageParser) getPriceEvent() response.WSPriceEvent {
	return response.WSPriceEvent{
		Event:         p.union.Event,
		Symbol:        p.union.Symbol,
		Currency:      p.union.Currency,
		Exchange:      p.union.Exchange,
		MicCode:       p.union.MicCode,
		Type:          p.union.Type,
		Timestamp:     p.union.Timestamp,
		Price:         p.union.Price,
		DayVolume:     p.union.DayVolume,
		Bid:           p.union.Bid,
		Ask:           p.union.Ask,
		CurrencyBase:  p.union.CurrencyBase,
		CurrencyQuote: p.union.CurrencyQuote,
	}
}

// getSubscribeStatusEvent converts the union to a WSSubscribeStatusEvent.
func (p *wsMessageParser) getSubscribeStatusEvent() response.WSSubscribeStatusEvent {
	return response.WSSubscribeStatusEvent{
		Event:   p.union.Event,
		Status:  p.union.Status,
		Success: p.union.Success,
		Fails:   p.union.Fails,
	}
}

// getErrorEvent converts the union to a WSErrorEvent.
func (p *wsMessageParser) getErrorEvent() response.WSErrorEvent {
	return response.WSErrorEvent{
		Event:   p.union.Event,
		Code:    p.union.Code,
		Message: p.union.Message,
		Status:  p.union.Status,
	}
}

// reset clears the cached union data, preparing the parser for reuse.
func (p *wsMessageParser) reset() {
	p.union = wsEventUnion{}
}
