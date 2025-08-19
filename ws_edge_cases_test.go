package twelvedata //nolint: testpackage

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/rs/zerolog"
)

// TestWS_CloseWithoutConnect tests closing WS client without connecting.
func TestWS_CloseWithoutConnect(t *testing.T) {
	t.Parallel()

	logger := zerolog.New(os.Stdout)
	cfg := &Conf{
		BaseWSURL: "ws.example.com",
		APIKey:    "test-key",
		WebSocket: WebSocket{PriceURL: "/test"},
	}

	ws := NewWS(cfg, &logger, nil)

	// Close without connecting should be safe
	err := ws.Close()
	if err != nil {
		t.Errorf("Close() without Connect() should not return error, got: %v", err)
	}

	// Multiple closes should be safe
	err = ws.Close()
	if err != nil {
		t.Errorf("Second Close() should not return error, got: %v", err)
	}
}

// TestWS_CloseBeforeMessageReaderStarts tests edge case where Close() is called immediately after Connect().
func TestWS_CloseBeforeMessageReaderStarts(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(_ *http.Request) bool { return true },
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Errorf("WebSocket upgrade failed: %v", err)
			return
		}
		defer func() { _ = conn.Close() }()

		// Keep connection open
		time.Sleep(1 * time.Second)
	}))
	defer server.Close()

	ws := createTestWS(t, server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := ws.Connect(ctx); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}

	// Close immediately - this should terminate goroutines cleanly
	if err := ws.Close(); err != nil {
		t.Errorf("Close() returned error: %v", err)
	}
}

// TestWS_OperationsAfterClose tests that operations after Close() handle gracefully.
func TestWS_OperationsAfterClose(t *testing.T) {
	t.Parallel()

	server := createMockWSServer(t, func(_ *websocket.Conn) {
		time.Sleep(500 * time.Millisecond)
	})
	defer server.Close()

	ws := createTestWS(t, server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := ws.Connect(ctx); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}

	// Close the WebSocket
	if err := ws.Close(); err != nil {
		t.Errorf("Close() returned error: %v", err)
	}

	// Try operations after close - they should return errors gracefully
	operations := []struct {
		name string
		op   func() error
	}{
		{"Subscribe", func() error { return ws.Subscribe([]string{"TEST"}) }},
		{"Unsubscribe", func() error { return ws.Unsubscribe([]string{"TEST"}) }},
		{"Reset", func() error { return ws.Reset() }},
		{"SendHeartbeat", func() error { return ws.SendHeartbeat() }},
	}

	for _, op := range operations {
		err := op.op()
		if err == nil {
			t.Errorf("%s should return error after Close(), but didn't", op.name)
		}
		t.Logf("%s after Close(): %v (expected)", op.name, err)
	}
}

// TestWS_ChannelConsumersAfterConnectionLoss tests channel behavior after connection loss.
func TestWS_ChannelConsumersAfterConnectionLoss(t *testing.T) {
	t.Parallel()

	var serverConn *websocket.Conn
	var serverMu sync.Mutex

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(_ *http.Request) bool { return true },
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Errorf("WebSocket upgrade failed: %v", err)
			return
		}
		defer func() { _ = conn.Close() }()

		serverMu.Lock()
		serverConn = conn
		serverMu.Unlock()

		// Send initial message
		msg := `{"event":"price","symbol":"LOSS","price":75.00,"timestamp":1643972766}`
		_ = conn.WriteMessage(websocket.TextMessage, []byte(msg))

		// Keep connection alive until externally closed
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				return
			}
		}
	}))
	defer server.Close()

	ws := createTestWS(t, server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := ws.Connect(ctx); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}

	// Start consumers
	priceConsumerBlocked := make(chan bool, 1)
	statusConsumerBlocked := make(chan bool, 1)
	errorConsumerBlocked := make(chan bool, 1)

	receivedEvents := make([]string, 0)
	var eventMu sync.Mutex

	go func() {
		for event := range ws.ConsumePriceEvents() {
			eventMu.Lock()
			receivedEvents = append(receivedEvents, "price")
			eventMu.Unlock()
			t.Logf("Received price event: %+v", event)
		}
		priceConsumerBlocked <- true
	}()

	go func() {
		for event := range ws.ConsumeStatusEvents() {
			eventMu.Lock()
			receivedEvents = append(receivedEvents, "status")
			eventMu.Unlock()
			t.Logf("Received status event: %+v", event)
		}
		statusConsumerBlocked <- true
	}()

	go func() {
		for event := range ws.ConsumeErrorEvents() {
			eventMu.Lock()
			receivedEvents = append(receivedEvents, "error")
			eventMu.Unlock()
			t.Logf("Received error event: %+v", event)
		}
		errorConsumerBlocked <- true
	}()

	// Wait for initial events
	time.Sleep(200 * time.Millisecond)

	// Force disconnect by closing server connection
	serverMu.Lock()
	if serverConn != nil {
		_ = serverConn.Close()
	}
	serverMu.Unlock()

	// Wait for disconnect to be detected
	time.Sleep(300 * time.Millisecond)

	// Close the client
	if err := ws.Close(); err != nil {
		t.Logf("Close() error after connection loss: %v", err)
	}

	// Check if consumers terminated
	checkConsumerTermination := func(ch <-chan bool, name string) {
		select {
		case <-ch:
			t.Logf("%s consumer terminated successfully", name)
		case <-time.After(500 * time.Millisecond):
			t.Errorf("%s consumer is still blocked after connection loss and Close()", name)
		}
	}

	checkConsumerTermination(priceConsumerBlocked, "Price")
	checkConsumerTermination(statusConsumerBlocked, "Status")
	checkConsumerTermination(errorConsumerBlocked, "Error")

	eventMu.Lock()
	t.Logf("Total events received before termination: %v", receivedEvents)
	eventMu.Unlock()
}

// TestWS_HighFrequencyConnectDisconnect tests rapid connect/disconnect cycles.
func TestWS_HighFrequencyConnectDisconnect(t *testing.T) {
	t.Parallel()

	server := createMockWSServer(t, func(conn *websocket.Conn) {
		// Send one message then close
		msg := `{"event":"price","symbol":"RAPID","price":25.00,"timestamp":1643972766}`
		_ = conn.WriteMessage(websocket.TextMessage, []byte(msg))
		time.Sleep(50 * time.Millisecond)
	})
	defer server.Close()

	initialGoroutines := runtime.NumGoroutine()

	// Rapid connect/disconnect cycles
	for i := 0; i < 10; i++ {
		ws := createTestWS(t, server.URL)

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

		if err := ws.Connect(ctx); err != nil {
			t.Logf("Connect #%d failed (may be expected): %v", i, err)
			cancel()
			continue
		}

		// Brief activity
		time.Sleep(10 * time.Millisecond)

		// Close
		if err := ws.Close(); err != nil {
			t.Logf("Close #%d error: %v", i, err)
		}

		cancel()
	}

	// Check for excessive goroutine growth
	time.Sleep(200 * time.Millisecond)
	runtime.GC()
	finalGoroutines := runtime.NumGoroutine()

	if finalGoroutines > initialGoroutines+5 {
		t.Errorf("Possible goroutine leak in rapid connect/disconnect: initial=%d, final=%d",
			initialGoroutines, finalGoroutines)
	}
}

// TestWS_ConcurrentChannelConsumption tests multiple concurrent consumers per channel.
func TestWS_ConcurrentChannelConsumption(t *testing.T) {
	t.Parallel()

	messageCount := 100

	server := createMockWSServer(t, func(conn *websocket.Conn) {
		for i := 0; i < messageCount; i++ {
			msg := `{"event":"price","symbol":"CONCURRENT","price":10.00,"timestamp":1643972766}`
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				return
			}
			time.Sleep(5 * time.Millisecond)
		}
		time.Sleep(100 * time.Millisecond)
	})
	defer server.Close()

	ws := createTestWS(t, server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ws.Connect(ctx); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer func() { _ = ws.Close() }()

	// Start multiple consumers for the same channel
	var wg sync.WaitGroup
	consumerCount := 5
	receivedCounts := make([]int, consumerCount)

	for i := 0; i < consumerCount; i++ {
		wg.Add(1)
		go func(consumerID int) {
			defer wg.Done()
			for range ws.ConsumePriceEvents() {
				receivedCounts[consumerID]++
			}
		}(i)
	}

	// Wait for messages to be sent
	time.Sleep(1 * time.Second)

	// Close WebSocket
	if err := ws.Close(); err != nil {
		t.Errorf("Close() error: %v", err)
	}

	// Wait for consumers to terminate (or timeout)
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		t.Log("All concurrent consumers terminated successfully")
		totalReceived := 0
		for i, count := range receivedCounts {
			totalReceived += count
			t.Logf("Consumer %d received %d events", i, count)
		}
		t.Logf("Total events received across all consumers: %d/%d", totalReceived, messageCount)
	case <-time.After(1 * time.Second):
		t.Error("Some concurrent consumers did not terminate after Close()")
	}
}

// TestWS_SendJSONMessage_TOCTOU_RaceCondition tests for the Time-of-Check-to-Time-of-Use race condition
// in sendJSONMessage where IsClosed() is checked before acquiring the lock.
func TestWS_SendJSONMessage_TOCTOU_RaceCondition(t *testing.T) {
	t.Parallel()

	// Create mock WebSocket server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(_ *http.Request) bool { return true },
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Errorf("WebSocket upgrade failed: %v", err)
			return
		}
		defer func() { _ = conn.Close() }()

		// Keep the connection alive during the test
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				return // Connection closed
			}
		}
	}))
	defer server.Close()

	logger := zerolog.New(os.Stdout).Level(zerolog.WarnLevel)
	cfg := &Conf{
		BaseWSURL: strings.TrimPrefix(strings.Replace(server.URL, "http://", "ws://", 1), "ws://"),
		APIKey:    "test-key",
		WebSocket: WebSocket{PriceURL: "/test"},
	}

	ws := NewWS(cfg, &logger, nil)

	// Override scheme to use non-TLS WebSocket for testing
	ws.url.Scheme = "ws"

	// Connect first
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ws.Connect(ctx); err != nil {
		t.Fatalf("Connect() failed: %v", err)
	}

	// Verify we're connected
	if ws.IsClosed() {
		t.Fatal("WebSocket should be connected, not closed")
	}

	const numGoroutines = 50
	const numMessages = 100

	var (
		wg           sync.WaitGroup
		sendErrors   = make(chan error, numGoroutines*numMessages)
		closeError   = make(chan error, 1)
		raceDetected = make(chan bool, 1)
	)

	// Create a test message to send
	testMessage := map[string]interface{}{
		"action": "subscribe",
		"params": map[string]string{"symbols": "AAPL"},
	}

	// Start goroutines that continuously send messages
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(_ int) {
			defer wg.Done()
			for j := 0; j < numMessages; j++ {
				if err := ws.sendJSONMessage(testMessage); err != nil {
					select {
					case sendErrors <- err:
					default:
						// Channel full, ignore
					}
				}
				// Small delay to increase chance of hitting race condition
				runtime.Gosched()
			}
		}(i)
	}

	// Wait a tiny bit to let senders start
	time.Sleep(10 * time.Millisecond)

	// Start a goroutine to close the connection while sends are happening
	go func() {
		// Multiple rapid close attempts to increase race window
		for i := 0; i < 10; i++ {
			if err := ws.Close(); err != nil {
				select {
				case closeError <- err:
					return
				default:
				}
			}
			time.Sleep(time.Microsecond) // Very brief delay
		}
		closeError <- nil
	}()

	// Wait for all senders to finish
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// Monitor for race conditions - when concurrent access happens,
	// we might see panics or unexpected errors
	go func() {
		defer func() {
			if r := recover(); r != nil {
				select {
				case raceDetected <- true:
				default:
				}
			}
		}()
		<-done
		raceDetected <- false
	}()

	select {
	case <-done:
		t.Log("All send operations completed")
	case <-time.After(5 * time.Second):
		t.Fatal("Test timed out - potential deadlock or hanging goroutines")
	}

	// Check results
	select {
	case err := <-closeError:
		if err != nil {
			t.Errorf("Close() returned unexpected error: %v", err)
		}
	case <-time.After(100 * time.Millisecond):
		// Close completed without reporting
	}

	select {
	case raceDetected := <-raceDetected:
		if raceDetected {
			t.Error("Race condition detected - panic occurred during concurrent operations")
		}
	default:
	}

	// Count and analyze send errors
	close(sendErrors)
	var errorCount int
	var closedErrors int
	for err := range sendErrors {
		errorCount++
		if err != nil && (err.Error() == "WebSocket is closed or not connected") {
			closedErrors++
		}
	}

	t.Logf("Total send errors: %d, closed-related errors: %d", errorCount, closedErrors)

	// The race condition we're testing for is:
	// 1. sendJSONMessage checks IsClosed() -> returns false
	// 2. Close() is called concurrently, sets closed=true
	// 3. sendJSONMessage acquires lock and continues with potentially invalid connection

	// With race detector enabled, this test should catch the race if it exists
}

// TestWS_SendJSONMessage_AfterClose verifies behavior when sending after close.
func TestWS_SendJSONMessage_AfterClose(t *testing.T) {
	t.Parallel()

	logger := zerolog.New(os.Stdout).Level(zerolog.WarnLevel)
	cfg := &Conf{
		BaseWSURL: "ws://example.com",
		APIKey:    "test-key",
		WebSocket: WebSocket{PriceURL: "/test"},
	}

	ws := NewWS(cfg, &logger, nil)

	// Close without connecting
	if err := ws.Close(); err != nil {
		t.Fatalf("Close() failed: %v", err)
	}

	// Try to send after close
	testMessage := map[string]interface{}{
		"action": "test",
	}

	err := ws.sendJSONMessage(testMessage)
	if err == nil {
		t.Error("sendJSONMessage() should return error when WebSocket is closed")
	}

	expectedErrMsg := "WebSocket is closed or not connected"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message '%s', got: '%s'", expectedErrMsg, err.Error())
	}
}
