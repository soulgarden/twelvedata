package twelvedata

import (
	"testing"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/response"
)

func TestWSMessageParser_parseMessage(t *testing.T) {
	tests := []struct {
		name          string
		message       []byte
		expectedEvent response.WSEventType
		shouldError   bool
	}{
		{
			name:          "price event",
			message:       []byte(`{"event":"price","symbol":"AAPL","price":150.25}`),
			expectedEvent: response.WSEventPrice,
			shouldError:   false,
		},
		{
			name:          "subscribe-status event",
			message:       []byte(`{"event":"subscribe-status","status":"ok"}`),
			expectedEvent: response.WSEventSubscribeStatus,
			shouldError:   false,
		},
		{
			name:          "error event",
			message:       []byte(`{"event":"error","message":"test error"}`),
			expectedEvent: response.WSEventError,
			shouldError:   false,
		},
		{
			name:        "malformed JSON",
			message:     []byte(`{"event":"price","symbol":}`),
			shouldError: true,
		},
		{
			name:        "missing event field",
			message:     []byte(`{"symbol":"AAPL","price":150.25}`),
			shouldError: false, // Should parse but return empty event type
		},
		{
			name:        "empty message",
			message:     []byte(`{}`),
			shouldError: false, // Should parse but return empty event type
		},
		{
			name:        "invalid JSON",
			message:     []byte(`not json`),
			shouldError: true,
		},
	}

	parser := newWSMessageParser()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eventType, err := parser.parseMessage(tt.message)

			if tt.shouldError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				if parser.union.Event != "" {
					t.Errorf("expected parser to reset on error, got event %q", parser.union.Event)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if eventType != tt.expectedEvent {
				t.Errorf("expected event type %q, got %q", tt.expectedEvent, eventType)
			}

			parser.reset()
		})
	}
}

func TestWSMessageParser_parsePrice(t *testing.T) {
	tests := []struct {
		name        string
		message     []byte
		expected    response.WSPriceEvent
		shouldError bool
	}{
		{
			name:    "basic price event",
			message: []byte(`{"event":"price","symbol":"AAPL","currency":"USD","exchange":"NASDAQ","type":"Common Stock","timestamp":1643972766,"price":150.25}`),
			expected: response.WSPriceEvent{
				Event:     response.WSEventPrice,
				Symbol:    "AAPL",
				Currency:  "USD",
				Exchange:  "NASDAQ",
				Type:      "Common Stock",
				Timestamp: 1643972766,
				Price:     150.25,
			},
			shouldError: false,
		},
		{
			name:    "forex price event with bid/ask",
			message: []byte(`{"event":"price","symbol":"EUR/USD","currency":"USD","exchange":"FOREX","type":"Currency","timestamp":1643972766,"price":1.1245,"bid":1.1244,"ask":1.1246,"currency_base":"EUR","currency_quote":"USD"}`),
			expected: response.WSPriceEvent{
				Event:         response.WSEventPrice,
				Symbol:        "EUR/USD",
				Currency:      "USD",
				Exchange:      "FOREX",
				Type:          "Currency",
				Timestamp:     1643972766,
				Price:         1.1245,
				Bid:           null.FloatFrom(1.1244),
				Ask:           null.FloatFrom(1.1246),
				CurrencyBase:  null.StringFrom("EUR"),
				CurrencyQuote: null.StringFrom("USD"),
			},
			shouldError: false,
		},
		{
			name:        "invalid JSON",
			message:     []byte(`{"event":"price","symbol":}`),
			shouldError: true,
		},
	}

	parser := newWSMessageParser()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the message in single pass
			_, err := parser.parseMessage(tt.message)
			if err != nil && !tt.shouldError {
				t.Fatalf("failed to parse message: %v", err)
			}

			priceEvent := parser.getPriceEvent()

			if tt.shouldError {
				parser.reset()
				return
			}

			// Compare basic fields
			if priceEvent.Event != tt.expected.Event {
				t.Errorf("expected event %q, got %q", tt.expected.Event, priceEvent.Event)
			}
			if priceEvent.Symbol != tt.expected.Symbol {
				t.Errorf("expected symbol %q, got %q", tt.expected.Symbol, priceEvent.Symbol)
			}
			if priceEvent.Price != tt.expected.Price {
				t.Errorf("expected price %f, got %f", tt.expected.Price, priceEvent.Price)
			}

			// Compare nullable fields if they exist
			if tt.expected.Bid.Valid && priceEvent.Bid.Float64 != tt.expected.Bid.Float64 {
				t.Errorf("expected bid %f, got %f", tt.expected.Bid.Float64, priceEvent.Bid.Float64)
			}

			parser.reset()
		})
	}
}

func TestWSMessageParser_parseSubscribeStatus(t *testing.T) {
	tests := []struct {
		name        string
		message     []byte
		expected    response.WSSubscribeStatusEvent
		shouldError bool
	}{
		{
			name:    "successful subscription",
			message: []byte(`{"event":"subscribe-status","status":"ok","success":[{"symbol":"AAPL","exchange":"NASDAQ","country":"United States","type":"Common Stock"}]}`),
			expected: response.WSSubscribeStatusEvent{
				Event:  response.WSEventSubscribeStatus,
				Status: "ok",
				Success: []response.WSSubscriptionResult{
					{
						Symbol:   "AAPL",
						Exchange: "NASDAQ",
						Country:  "United States",
						Type:     "Common Stock",
					},
				},
			},
			shouldError: false,
		},
		{
			name:    "failed subscription",
			message: []byte(`{"event":"subscribe-status","status":"error","fails":[{"symbol":"INVALID","message":"Symbol not found"}]}`),
			expected: response.WSSubscribeStatusEvent{
				Event:  response.WSEventSubscribeStatus,
				Status: "error",
				Fails: []response.WSSubscriptionFail{
					{
						Symbol:  "INVALID",
						Message: "Symbol not found",
					},
				},
			},
			shouldError: false,
		},
	}

	parser := newWSMessageParser()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the message in single pass
			_, err := parser.parseMessage(tt.message)
			if err != nil {
				t.Fatalf("failed to parse message: %v", err)
			}

			statusEvent := parser.getSubscribeStatusEvent()

			if tt.shouldError {
				parser.reset()
				return
			}

			if statusEvent.Event != tt.expected.Event {
				t.Errorf("expected event %q, got %q", tt.expected.Event, statusEvent.Event)
			}
			if statusEvent.Status != tt.expected.Status {
				t.Errorf("expected status %q, got %q", tt.expected.Status, statusEvent.Status)
			}

			parser.reset()
		})
	}
}

func TestWSMessageParser_parseError(t *testing.T) {
	tests := []struct {
		name        string
		message     []byte
		expected    response.WSErrorEvent
		shouldError bool
	}{
		{
			name:    "basic error event",
			message: []byte(`{"event":"error","code":"400","message":"Invalid symbol","status":"error"}`),
			expected: response.WSErrorEvent{
				Event:   response.WSEventError,
				Code:    "400",
				Message: "Invalid symbol",
				Status:  "error",
			},
			shouldError: false,
		},
		{
			name:    "error event without optional fields",
			message: []byte(`{"event":"error","message":"Something went wrong"}`),
			expected: response.WSErrorEvent{
				Event:   response.WSEventError,
				Message: "Something went wrong",
			},
			shouldError: false,
		},
	}

	parser := newWSMessageParser()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the message in single pass
			_, err := parser.parseMessage(tt.message)
			if err != nil {
				t.Fatalf("failed to parse message: %v", err)
			}

			errorEvent := parser.getErrorEvent()

			if tt.shouldError {
				parser.reset()
				return
			}

			if errorEvent.Event != tt.expected.Event {
				t.Errorf("expected event %q, got %q", tt.expected.Event, errorEvent.Event)
			}
			if errorEvent.Message != tt.expected.Message {
				t.Errorf("expected message %q, got %q", tt.expected.Message, errorEvent.Message)
			}
			if errorEvent.Code != tt.expected.Code {
				t.Errorf("expected code %q, got %q", tt.expected.Code, errorEvent.Code)
			}

			parser.reset()
		})
	}
}

func TestWSMessageParser_reset(t *testing.T) {
	parser := newWSMessageParser()
	message := []byte(`{"event":"price","symbol":"AAPL","price":150.25}`)

	// Parse message to populate union
	_, err := parser.parseMessage(message)
	if err != nil {
		t.Fatalf("failed to parse message: %v", err)
	}

	// Verify data is cached in union
	if parser.union.Event == "" {
		t.Error("expected union to contain parsed data")
	}

	// Reset parser
	parser.reset()

	// Verify data is cleared
	if parser.union.Event != "" {
		t.Error("expected union data to be cleared after reset")
	}
}
