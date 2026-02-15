// Package main demonstrates comprehensive error handling patterns for the Twelve Data API client.
package main

import (
	"fmt"
	"log"

	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata"
	"github.com/soulgarden/twelvedata/request"
	"github.com/valyala/fasthttp"
)

func main() {
	logger := zerolog.New(nil).Level(zerolog.Disabled) // Disable logging for example

	cfg := &twelvedata.Conf{
		APIKey: "invalid-api-key", // This will trigger a 401 error
	}

	httpCli := twelvedata.NewHTTPCli(&fasthttp.Client{}, cfg, &logger)
	client := twelvedata.NewClient(httpCli, cfg)

	// Try to get stocks with invalid API key
	_, _, err := client.GetStocks(request.GetStock{
		APIKey: request.APIKey{
			APIKey: cfg.APIKey,
		},
	})

	if err != nil {
		fmt.Printf("Got error: %v\n", err)

		// Check error type and handle accordingly
		switch {
		// Domain-specific errors (these are parsed from API response messages)
		case twelvedata.IsSymbolNotFoundError(err):
			fmt.Println("❌ Symbol not found")

		case twelvedata.IsPlanLimitationError(err):
			fmt.Println("❌ Plan limitation - feature not available in your plan")

		case twelvedata.IsInsufficientCreditsError(err):
			fmt.Println("❌ Insufficient API credits")

		case twelvedata.IsAPIKeyError(err):
			fmt.Println("❌ API key issue")

		// HTTP-level errors
		case twelvedata.IsUnauthorizedError(err):
			fmt.Println("❌ Authentication failed - check your API key")

		case twelvedata.IsBadRequestError(err):
			fmt.Println("❌ Bad request - check your parameters")

		case twelvedata.IsRateLimitError(err):
			fmt.Println("❌ Rate limit exceeded - wait and retry")

		case twelvedata.IsNotFoundError(err):
			fmt.Println("❌ Resource not found")

		case twelvedata.IsTimeoutError(err):
			fmt.Println("❌ Request timeout - check network connection")

		case twelvedata.IsNetworkError(err):
			fmt.Println("❌ Network error - check connectivity")

		// WebSocket-specific errors
		case twelvedata.IsWSConnectionError(err):
			fmt.Println("❌ WebSocket connection error")

		case twelvedata.IsWSMessageError(err):
			fmt.Println("❌ WebSocket message error")

		case twelvedata.IsWSSubscriptionError(err):
			fmt.Println("❌ WebSocket subscription error")

		// Generic error categories
		case twelvedata.IsDomainError(err):
			fmt.Println("❌ Domain error - business logic issue")

		case twelvedata.IsWSError(err):
			fmt.Println("❌ WebSocket error")

		case twelvedata.IsHTTPError(err):
			fmt.Println("❌ HTTP error - server responded with error")

		default:
			fmt.Println("❌ Unknown error occurred")
		}
	}

	fmt.Println("\n🔄 Example with successful request:")

	// Example with valid configuration
	validCfg := &twelvedata.Conf{
		APIKey: "demo", // Demo API key should work
	}

	validHTTPCli := twelvedata.NewHTTPCli(&fasthttp.Client{}, validCfg, &logger)
	validClient := twelvedata.NewClient(validHTTPCli, validCfg)

	stocks, credits, err := validClient.GetStocks(request.GetStock{
		APIKey: request.APIKey{
			APIKey: validCfg.APIKey,
		},
	})

	if err != nil {
		log.Printf("Unexpected error with demo key: %v", err)
	} else {
		fmt.Printf("✅ Successfully retrieved %d stocks\n", len(stocks.Data))
		fmt.Printf("📊 Credits used: %d, remaining: %d\n",
			credits.GetCreditsUsed(), credits.GetCreditsLeft())
	}
}
