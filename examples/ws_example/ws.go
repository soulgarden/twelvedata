// Package main demonstrates WebSocket usage for real-time market data streaming.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/jinzhu/configor"
	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata"
	"github.com/soulgarden/twelvedata/request"
)

func main() {
	// Setup structured logging
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	// Load configuration
	cfg := &twelvedata.Conf{
		APIKey: "demo", // Set API key
	}

	if err := configor.New(&configor.Config{}).Load(cfg); err != nil {
		logger.Fatal().Err(err).Msg("failed to load config")
	}

	// Create WebSocket client
	wsCli := twelvedata.NewWS(cfg, &logger, nil)

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	// Setup signal handling for graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// WaitGroup to track consumer goroutines
	var wg sync.WaitGroup

	// Connect to WebSocket
	if err := wsCli.Connect(ctx); err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to WebSocket")
	}

	logger.Info().Msg("WebSocket connected successfully")

	// Start consuming events in separate goroutines with context awareness
	wg.Add(3)
	go consumePriceEvents(ctx, wsCli, &wg, &logger)
	go consumeStatusEvents(ctx, wsCli, &wg, &logger)
	go consumeErrorEvents(ctx, wsCli, &wg, &logger)

	// Example 1: Simple subscription
	simpleSymbols := []string{"AAPL", "MSFT", "GOOGL"}
	if err := wsCli.Subscribe(simpleSymbols); err != nil {
		logger.Error().Err(err).Msg("failed to subscribe to simple symbols")
	} else {
		logger.Info().Strs("symbols", simpleSymbols).Msg("subscribed to simple symbols")
	}

	// Wait a bit then demonstrate extended subscription
	time.Sleep(2 * time.Second)

	// Example 2: Extended subscription format
	extendedSymbols := []request.WSSymbolExtended{
		{
			Symbol:   "TSLA",
			Exchange: "NASDAQ",
			Type:     "Common Stock",
		},
		{
			Symbol: "EUR/USD",
			Type:   "Forex",
		},
	}
	if err := wsCli.SubscribeExtended(extendedSymbols); err != nil {
		logger.Error().Err(err).Msg("failed to subscribe to extended symbols")
	} else {
		logger.Info().Int("count", len(extendedSymbols)).Msg("subscribed to extended symbols")
	}

	// Wait for events to come in
	time.Sleep(5 * time.Second)

	// Example 3: Demonstrate unsubscribe
	if err := wsCli.Unsubscribe([]string{"MSFT"}); err != nil {
		logger.Error().Err(err).Msg("failed to unsubscribe from MSFT")
	} else {
		logger.Info().Msg("unsubscribed from MSFT")
	}

	// Example 4: Send manual heartbeat
	if err := wsCli.SendHeartbeat(); err != nil {
		logger.Error().Err(err).Msg("failed to send heartbeat")
	} else {
		logger.Info().Msg("manual heartbeat sent")
	}

	// Wait for more events
	time.Sleep(5 * time.Second)

	// Example 5: Reset all subscriptions
	if err := wsCli.Reset(); err != nil {
		logger.Error().Err(err).Msg("failed to reset subscriptions")
	} else {
		logger.Info().Msg("reset all subscriptions")
	}

	// Wait for termination signal
	<-sigCh
	logger.Info().Msg("shutdown signal received, initiating graceful shutdown...")

	// Step 1: Cancel context to signal all goroutines to stop
	cancel()
	logger.Info().Msg("context cancelled, waiting for consumer goroutines to finish...")

	// Step 2: Wait for consumer goroutines to finish gracefully (with timeout)
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		logger.Info().Msg("all consumer goroutines stopped gracefully")
	case <-time.After(5 * time.Second):
		logger.Warn().Msg("timeout waiting for consumer goroutines, continuing with shutdown")
	}

	// Step 3: Close WebSocket connection
	logger.Info().Msg("closing WebSocket connection...")
	if err := wsCli.Close(); err != nil {
		logger.Error().Err(err).Msg("error closing WebSocket connection")
	} else {
		logger.Info().Msg("WebSocket connection closed successfully")
	}

	logger.Info().Msg("graceful shutdown completed")
}

func consumePriceEvents(ctx context.Context, wsCli *twelvedata.WS, wg *sync.WaitGroup, logger *zerolog.Logger) {
	defer wg.Done()
	defer logger.Debug().Msg("price events consumer stopped")

	for {
		select {
		case <-ctx.Done():
			logger.Debug().Msg("price events consumer received shutdown signal")
			return
		case event, ok := <-wsCli.ConsumePriceEvents():
			if !ok {
				logger.Debug().Msg("price events channel closed")
				return
			}
			price := "n/a"
			if event.Price.Valid {
				price = fmt.Sprintf("%.4f", event.Price.Float64)
			}
			timestamp := "n/a"
			if event.Timestamp.Valid {
				timestamp = fmt.Sprintf("%d", event.Timestamp.Int64)
			}

			fmt.Printf("🔥 Price Event: %s @ $%s (%s) [%s]\n",
				event.Symbol, price, event.Exchange, timestamp)

			// Show additional fields for forex/crypto if available
			if event.Bid.Valid || event.Ask.Valid {
				fmt.Printf("   📊 Bid: %.4f, Ask: %.4f\n",
					event.Bid.Float64, event.Ask.Float64)
			}

			if event.DayVolume.Valid {
				fmt.Printf("   📈 Volume: %d\n", event.DayVolume.Int64)
			}
		}
	}
}

func consumeStatusEvents(ctx context.Context, wsCli *twelvedata.WS, wg *sync.WaitGroup, logger *zerolog.Logger) {
	defer wg.Done()
	defer logger.Debug().Msg("status events consumer stopped")

	for {
		select {
		case <-ctx.Done():
			logger.Debug().Msg("status events consumer received shutdown signal")
			return
		case event, ok := <-wsCli.ConsumeStatusEvents():
			if !ok {
				logger.Debug().Msg("status events channel closed")
				return
			}
			fmt.Printf("ℹ️  Status: %s\n", event.Status)

			if len(event.Success) > 0 {
				fmt.Printf("   ✅ Successfully subscribed to %d symbols:\n", len(event.Success))
				for _, success := range event.Success {
					fmt.Printf("      - %s (%s, %s)\n",
						success.Symbol, success.Exchange, success.Type)
				}
			}

			if len(event.Fails) > 0 {
				fmt.Printf("   ❌ Failed to subscribe to %d symbols:\n", len(event.Fails))
				for _, fail := range event.Fails {
					fmt.Printf("      - %s: %s\n", fail.Symbol, fail.Message)
				}
			}
		}
	}
}

func consumeErrorEvents(ctx context.Context, wsCli *twelvedata.WS, wg *sync.WaitGroup, logger *zerolog.Logger) {
	defer wg.Done()
	defer logger.Debug().Msg("error events consumer stopped")

	for {
		select {
		case <-ctx.Done():
			logger.Debug().Msg("error events consumer received shutdown signal")
			return
		case event, ok := <-wsCli.ConsumeErrorEvents():
			if !ok {
				logger.Debug().Msg("error events channel closed")
				return
			}
			log.Printf("❌ Error Event: %s - %s\n", event.Code, event.Message)
		}
	}
}
