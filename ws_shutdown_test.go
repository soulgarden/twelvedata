package twelvedata //nolint: testpackage

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/rs/zerolog"
)

// TestWS_GracefulShutdown_ChannelsClosed verifies that all channels are properly closed on shutdown.
func TestWS_GracefulShutdown_ChannelsClosed(t *testing.T) {
	t.Parallel()

	server := createMockWSServer(t, func(_ *websocket.Conn) {
		// Keep connection alive for test duration
		time.Sleep(500 * time.Millisecond)
	})
	defer server.Close()

	ws := createTestWS(t, server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Connect to server
	err := ws.Connect(ctx)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}

	// Start consumers that will detect if channels are closed
	priceChannelClosed := make(chan bool, 1)
	statusChannelClosed := make(chan bool, 1)
	errorChannelClosed := make(chan bool, 1)

	// Consumer for price events
	go func() {
		//nolint:revive // Intentionally draining channel
		for range ws.ConsumePriceEvents() {
			// Drain channel events until closed
		}
		priceChannelClosed <- true
	}()

	// Consumer for status events
	go func() {
		//nolint:revive // Intentionally draining channel
		for range ws.ConsumeStatusEvents() {
			// Drain channel events until closed
		}
		statusChannelClosed <- true
	}()

	// Consumer for error events
	go func() {
		//nolint:revive // Intentionally draining channel
		for range ws.ConsumeErrorEvents() {
			// Drain channel events until closed
		}
		errorChannelClosed <- true
	}()

	// Close WebSocket
	if err := ws.Close(); err != nil {
		t.Errorf("Close() returned error: %v", err)
	}

	// Check if channels were closed (they should be, but currently aren't)
	checkChannelClosed := func(ch <-chan bool, name string) {
		select {
		case closed := <-ch:
			if !closed {
				t.Errorf("%s channel consumer did not detect channel closure", name)
			}
		case <-time.After(100 * time.Millisecond):
			t.Errorf("%s channel was NOT closed after shutdown (consumer still blocked)", name)
		}
	}

	checkChannelClosed(priceChannelClosed, "Price")
	checkChannelClosed(statusChannelClosed, "Status")
	checkChannelClosed(errorChannelClosed, "Error")
}

// TestWS_GracefulShutdown_GoroutineLeaks checks for goroutine leaks after closing.
func TestWS_GracefulShutdown_GoroutineLeaks(t *testing.T) {
	t.Parallel()

	// Get initial goroutine count
	runtime.GC()
	initialGoroutines := runtime.NumGoroutine()

	server := createMockWSServer(t, func(conn *websocket.Conn) {
		// Send periodic messages to keep reader active
		for i := 0; i < 5; i++ {
			msg := `{"event":"price","symbol":"TEST","price":100.00,"timestamp":1643972766}`
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				return
			}
			time.Sleep(50 * time.Millisecond)
		}
	})
	defer server.Close()

	// Run multiple connect/disconnect cycles to detect leaks
	for i := 0; i < 3; i++ {
		ws := createTestWS(t, server.URL)

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

		if err := ws.Connect(ctx); err != nil {
			t.Fatalf("Iteration %d: Failed to connect: %v", i, err)
		}

		// Let it run for a bit
		time.Sleep(100 * time.Millisecond)

		// Close connection
		if err := ws.Close(); err != nil {
			t.Errorf("Iteration %d: Close() error: %v", i, err)
		}

		cancel()
	}

	// Give time for goroutines to terminate
	time.Sleep(200 * time.Millisecond)
	runtime.GC()

	// Check final goroutine count
	finalGoroutines := runtime.NumGoroutine()
	leaked := finalGoroutines - initialGoroutines

	// Allow for some tolerance (test framework itself might create goroutines)
	if leaked > 2 {
		t.Errorf("Goroutine leak detected: %d goroutines leaked (initial: %d, final: %d)",
			leaked, initialGoroutines, finalGoroutines)

		// Print stack traces for debugging
		buf := make([]byte, 1<<16)
		stackLen := runtime.Stack(buf, true)
		t.Logf("Current goroutine stack traces:\n%s", buf[:stackLen])
	}
}

// TestWS_GracefulShutdown_RaceConditions tests for race conditions during shutdown.
func TestWS_GracefulShutdown_RaceConditions(t *testing.T) {
	t.Parallel()

	var messagesSent atomic.Int32
	var messagesReceived atomic.Int32

	server := createMockWSServer(t, func(conn *websocket.Conn) {
		// Continuously send messages
		for {
			msg := `{"event":"price","symbol":"RACE","price":99.99,"timestamp":1643972766}`
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				return
			}
			messagesSent.Add(1)
			time.Sleep(10 * time.Millisecond)
		}
	})
	defer server.Close()

	ws := createTestWS(t, server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ws.Connect(ctx); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}

	// Start multiple consumers
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range ws.ConsumePriceEvents() {
				messagesReceived.Add(1)
			}
		}()
	}

	// Let it run
	time.Sleep(200 * time.Millisecond)

	// Try to close while messages are being processed
	closeDone := make(chan error, 1)
	go func() {
		closeDone <- ws.Close()
	}()

	// Try to send operations while closing (should handle gracefully)
	go func() {
		_ = ws.SendHeartbeat()
		_ = ws.Subscribe([]string{"TEST"})
		_ = ws.Reset()
	}()

	// Wait for close to complete
	select {
	case err := <-closeDone:
		if err != nil {
			t.Errorf("Close() returned error during race test: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("Close() did not complete within timeout during race test")
	}

	// Verify no panic occurred and goroutines terminated
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Good, all consumers terminated
	case <-time.After(500 * time.Millisecond):
		t.Error("Consumer goroutines did not terminate after Close()")
	}

	t.Logf("Messages sent: %d, received: %d", messagesSent.Load(), messagesReceived.Load())
}

// TestWS_ConsumerBlockingAfterDisconnect verifies consumers don't block forever after disconnect.
func TestWS_ConsumerBlockingAfterDisconnect(t *testing.T) {
	t.Parallel()

	var serverConn *websocket.Conn
	var serverConnMu sync.Mutex

	server := createMockWSServer(t, func(conn *websocket.Conn) {
		serverConnMu.Lock()
		serverConn = conn
		serverConnMu.Unlock()

		// Send initial message
		msg := `{"event":"price","symbol":"BLOCK","price":50.00,"timestamp":1643972766}`
		_ = conn.WriteMessage(websocket.TextMessage, []byte(msg))

		// Keep connection open
		time.Sleep(1 * time.Second)
	})
	defer server.Close()

	ws := createTestWS(t, server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := ws.Connect(ctx); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}

	// Start a consumer
	consumerExited := make(chan bool, 1)
	receivedCount := atomic.Int32{}

	go func() {
		for event := range ws.ConsumePriceEvents() {
			receivedCount.Add(1)
			t.Logf("Received event: %+v", event)
		}
		consumerExited <- true
	}()

	// Wait for first message to be received
	time.Sleep(100 * time.Millisecond)

	// Force disconnect from server side
	serverConnMu.Lock()
	if serverConn != nil {
		_ = serverConn.Close()
	}
	serverConnMu.Unlock()

	// Give time for disconnect to be detected
	time.Sleep(200 * time.Millisecond)

	// Close the WebSocket client
	if err := ws.Close(); err != nil {
		t.Logf("Close() error (expected due to broken connection): %v", err)
	}

	// Check if consumer exited
	select {
	case exited := <-consumerExited:
		if !exited {
			t.Error("Consumer did not exit properly")
		}
		t.Logf("Consumer exited successfully after receiving %d events", receivedCount.Load())
	case <-time.After(500 * time.Millisecond):
		t.Error("Consumer is still blocked after connection was closed")
	}
}

// TestWS_ErrorPropagationDuringShutdown tests that errors are properly propagated during shutdown.
func TestWS_ErrorPropagationDuringShutdown(t *testing.T) {
	t.Parallel()

	errorSent := atomic.Bool{}

	server := createMockWSServer(t, func(conn *websocket.Conn) {
		// Send an error event
		errorMsg := `{"event":"error","code":"test_error","message":"Connection will be terminated"}`
		if err := conn.WriteMessage(websocket.TextMessage, []byte(errorMsg)); err != nil {
			t.Logf("Failed to send error message: %v", err)
			return
		}
		errorSent.Store(true)

		// Then close the connection abruptly
		time.Sleep(50 * time.Millisecond)
		_ = conn.Close()
	})
	defer server.Close()

	ws := createTestWS(t, server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := ws.Connect(ctx); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}

	// Monitor error events
	errorReceived := atomic.Bool{}
	go func() {
		for event := range ws.ConsumeErrorEvents() {
			t.Logf("Received error event: %+v", event)
			errorReceived.Store(true)
		}
	}()

	// Wait for error to be processed
	time.Sleep(200 * time.Millisecond)

	// Close should handle the broken connection gracefully
	err := ws.Close()
	if err != nil {
		t.Logf("Close() returned error (may be expected): %v", err)
	}

	// Verify error was received before shutdown
	if errorSent.Load() && !errorReceived.Load() {
		t.Error("Error event was sent but not received by consumer")
	}
}

// TestWS_MultipleCloseIdempotent tests that calling Close() multiple times is safe.
func TestWS_MultipleCloseIdempotent(t *testing.T) {
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

	// Call Close() multiple times concurrently
	var wg sync.WaitGroup
	errors := make([]error, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			errors[idx] = ws.Close()
		}(i)
	}

	wg.Wait()

	// At least one should succeed, others might return "already closed" or similar
	successCount := 0
	for i, err := range errors {
		if err == nil {
			successCount++
		}
		t.Logf("Close() call %d: %v", i, err)
	}

	if successCount == 0 {
		t.Error("None of the Close() calls succeeded")
	}
}

// TestWS_ChannelBufferOverflow tests behavior when event channels overflow.
func TestWS_ChannelBufferOverflow(t *testing.T) {
	t.Parallel()

	server := createMockWSServer(t, func(conn *websocket.Conn) {
		// Send many messages quickly to overflow the buffer
		for i := 0; i < 2000; i++ {
			msg := `{"event":"price","symbol":"OVERFLOW","price":100.00,"timestamp":1643972766}`
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				return
			}
		}
		time.Sleep(100 * time.Millisecond)
	})
	defer server.Close()

	// Create WS with small buffer for testing
	logger := zerolog.New(os.Stdout).Level(zerolog.WarnLevel)
	cfg := &Conf{
		BaseWSURL: strings.TrimPrefix(strings.Replace(server.URL, "http://", "ws://", 1), "ws://"),
		APIKey:    "test-key",
		WebSocket: WebSocket{PriceURL: "/quotes/price"},
	}

	ws := NewWS(cfg, &logger, nil)
	ws.url.Scheme = "ws"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := ws.Connect(ctx); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer func() { _ = ws.Close() }()

	// Slow consumer to cause buffer overflow
	receivedCount := atomic.Int32{}
	go func() {
		for range ws.ConsumePriceEvents() {
			receivedCount.Add(1)
			time.Sleep(10 * time.Millisecond) // Slow processing
		}
	}()

	// Wait for messages to be sent
	time.Sleep(500 * time.Millisecond)

	// Close and check
	if err := ws.Close(); err != nil {
		t.Logf("Close() error: %v", err)
	}

	t.Logf("Messages received: %d out of 2000 sent (dropped: ~%d)",
		receivedCount.Load(), 2000-receivedCount.Load())

	// We expect some messages to be dropped due to buffer overflow
	if receivedCount.Load() >= 2000 {
		t.Error("All messages were received, buffer overflow handling was not tested")
	}
}

// Helper functions

func createMockWSServer(t *testing.T, handler func(*websocket.Conn)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(_ *http.Request) bool {
				return true
			},
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Errorf("WebSocket upgrade failed: %v", err)
			return
		}
		defer func() { _ = conn.Close() }()

		handler(conn)
	}))
}

func createTestWS(_ *testing.T, serverURL string) *WS {
	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel)
	wsURL := strings.Replace(serverURL, "http://", "ws://", 1)
	cfg := &Conf{
		BaseWSURL: strings.TrimPrefix(wsURL, "ws://"),
		APIKey:    "test-key",
		WebSocket: WebSocket{PriceURL: "/quotes/price"},
	}

	ws := NewWS(cfg, &logger, nil)
	ws.url.Scheme = "ws"
	return ws
}
