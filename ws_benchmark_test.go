package twelvedata

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/rs/zerolog"

	"github.com/soulgarden/twelvedata/response"
)

// Test message samples representing different event types.
var (
	priceMessage = []byte(`{"event":"price","symbol":"AAPL","currency":"USD","exchange":"NASDAQ","type":"Common Stock","timestamp":1643972766,"price":150.25,"day_volume":12345678}`)

	statusMessage = []byte(`{"event":"subscribe-status","status":"ok","success":[{"symbol":"AAPL","exchange":"NASDAQ","country":"United States","type":"Common Stock"}]}`)

	errorMessage = []byte(`{"event":"error","code":"400","message":"Invalid symbol","status":"error"}`)

	// Complex price message with all fields.
	complexPriceMessage = []byte(`{"event":"price","symbol":"EUR/USD","currency":"USD","exchange":"FOREX","mic_code":"FXCM","type":"Currency","timestamp":1643972766,"price":1.1245,"bid":1.1244,"ask":1.1246,"currency_base":"EUR","currency_quote":"USD"}`)
)

// BenchmarkOldRouteMessagePriceEvent simulates the old double-unmarshal approach.
func BenchmarkOldRouteMessagePriceEvent(b *testing.B) {
	message := priceMessage

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		// Simulate old approach: double unmarshal
		var baseEvent response.WSBaseEvent
		if err := json.Unmarshal(message, &baseEvent); err != nil {
			b.Fatal(err)
		}

		if baseEvent.Event == response.WSEventPrice {
			var priceEvent response.WSPriceEvent
			if err := json.Unmarshal(message, &priceEvent); err != nil {
				b.Fatal(err)
			}
		}
	}
}

// BenchmarkNewRouteMessagePriceEvent tests the optimized single-pass parser.
func BenchmarkNewRouteMessagePriceEvent(b *testing.B) {
	parser := newWSMessageParser()
	message := priceMessage

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		// New approach: single pass parsing with union type
		eventType, err := parser.parseMessage(message)
		if err != nil {
			b.Fatal(err)
		}

		if eventType == response.WSEventPrice {
			_ = parser.getPriceEvent()
		}

		parser.reset()
	}
}

// BenchmarkOldRouteMessageStatusEvent simulates the old approach for status events.
func BenchmarkOldRouteMessageStatusEvent(b *testing.B) {
	message := statusMessage

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var baseEvent response.WSBaseEvent
		if err := json.Unmarshal(message, &baseEvent); err != nil {
			b.Fatal(err)
		}

		if baseEvent.Event == response.WSEventSubscribeStatus {
			var statusEvent response.WSSubscribeStatusEvent
			if err := json.Unmarshal(message, &statusEvent); err != nil {
				b.Fatal(err)
			}
		}
	}
}

// BenchmarkNewRouteMessageStatusEvent tests the optimized parser for status events.
func BenchmarkNewRouteMessageStatusEvent(b *testing.B) {
	parser := newWSMessageParser()
	message := statusMessage

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		eventType, err := parser.parseMessage(message)
		if err != nil {
			b.Fatal(err)
		}

		if eventType == response.WSEventSubscribeStatus {
			_ = parser.getSubscribeStatusEvent()
		}

		parser.reset()
	}
}

// BenchmarkOldRouteMessageComplexPrice tests old approach with complex price data.
func BenchmarkOldRouteMessageComplexPrice(b *testing.B) {
	message := complexPriceMessage

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var baseEvent response.WSBaseEvent
		if err := json.Unmarshal(message, &baseEvent); err != nil {
			b.Fatal(err)
		}

		if baseEvent.Event == response.WSEventPrice {
			var priceEvent response.WSPriceEvent
			if err := json.Unmarshal(message, &priceEvent); err != nil {
				b.Fatal(err)
			}
		}
	}
}

// BenchmarkNewRouteMessageComplexPrice tests optimized approach with complex price data.
func BenchmarkNewRouteMessageComplexPrice(b *testing.B) {
	parser := newWSMessageParser()
	message := complexPriceMessage

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		eventType, err := parser.parseMessage(message)
		if err != nil {
			b.Fatal(err)
		}

		if eventType == response.WSEventPrice {
			_ = parser.getPriceEvent()
		}

		parser.reset()
	}
}

// BenchmarkFullWSRouteMessage tests the full WebSocket routeMessage method.
func BenchmarkFullWSRouteMessage(b *testing.B) {
	logger := zerolog.Nop()
	cfg := &Conf{
		BaseURL:   "https://api.twelvedata.com",
		BaseWSURL: "ws.twelvedata.com",
		APIKey:    "test",
		WebSocket: WebSocket{PriceURL: "/v1/quotes/price"},
	}

	ws := NewWS(cfg, &logger, nil)

	// Initialize required fields for routeMessage to work properly
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ws.ctx = ctx
	ws.shutdown = make(chan struct{})
	defer close(ws.shutdown)

	message := priceMessage

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ws.routeMessage(message)
	}
}

// BenchmarkMessageMix tests realistic mixed message processing.
func BenchmarkMessageMix(b *testing.B) {
	parser := newWSMessageParser()
	messages := [][]byte{priceMessage, statusMessage, errorMessage, complexPriceMessage}
	messageCount := len(messages)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		message := messages[i%messageCount]

		eventType, err := parser.parseMessage(message)
		if err != nil {
			b.Fatal(err)
		}

		switch eventType {
		case response.WSEventPrice:
			_ = parser.getPriceEvent()
		case response.WSEventSubscribeStatus:
			_ = parser.getSubscribeStatusEvent()
		case response.WSEventError:
			_ = parser.getErrorEvent()
		}

		parser.reset()
	}
}

// BenchmarkParserReuse tests parser reuse efficiency.
func BenchmarkParserReuse(b *testing.B) {
	parser := newWSMessageParser()
	message := priceMessage

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		eventType, err := parser.parseMessage(message)
		if err != nil {
			b.Fatal(err)
		}

		if eventType == response.WSEventPrice {
			_ = parser.getPriceEvent()
		}

		parser.reset()
	}
}

// BenchmarkWithContext tests performance with context checking (realistic usage).
func BenchmarkWithContext(b *testing.B) {
	parser := newWSMessageParser()
	message := priceMessage
	ctx := context.Background()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		// Simulate context checking like in real routeMessage
		select {
		case <-ctx.Done():
			return
		default:
		}

		eventType, err := parser.parseMessage(message)
		if err != nil {
			b.Fatal(err)
		}

		if eventType == response.WSEventPrice {
			_ = parser.getPriceEvent()
		}

		parser.reset()
	}
}
