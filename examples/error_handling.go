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
		ApiKey: request.ApiKey{
			ApiKey: cfg.APIKey,
		},
	})

	if err != nil {
		fmt.Printf("Got error: %v\n", err)

		// Check error type and handle accordingly
		switch {
		case twelvedata.IsUnauthorizedError(err):
			fmt.Println("❌ Authentication failed - check your API key")
			handleUnauthorizedError()

		case twelvedata.IsBadRequestError(err):
			fmt.Println("❌ Bad request - check your parameters")
			handleBadRequestError()

		case twelvedata.IsRateLimitError(err):
			fmt.Println("❌ Rate limit exceeded - wait and retry")
			handleRateLimitError()

		case twelvedata.IsNotFoundError(err):
			fmt.Println("❌ Resource not found")
			handleNotFoundError()

		case twelvedata.IsTimeoutError(err):
			fmt.Println("❌ Request timeout - check network connection")
			handleTimeoutError()

		case twelvedata.IsNetworkError(err):
			fmt.Println("❌ Network error - check connectivity")
			handleNetworkError()

		case twelvedata.IsHTTPError(err):
			fmt.Println("❌ HTTP error - server responded with error")
			handleHTTPError()

		default:
			fmt.Println("❌ Unknown error occurred")
			handleGenericError()
		}
	}

	fmt.Println("\n🔄 Example with successful request:")

	// Example with valid configuration
	validCfg := &twelvedata.Conf{
		APIKey: "demo", // Demo API key should work
	}

	validHttpCli := twelvedata.NewHTTPCli(&fasthttp.Client{}, validCfg, &logger)
	validClient := twelvedata.NewClient(validHttpCli, validCfg)

	stocks, credits, err := validClient.GetStocks(request.GetStock{
		ApiKey: request.ApiKey{
			ApiKey: validCfg.APIKey,
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

func handleUnauthorizedError() {
	fmt.Println("💡 Solutions:")
	fmt.Println("   - Verify your API key is correct")
	fmt.Println("   - Check if your API key has expired")
	fmt.Println("   - Visit https://twelvedata.com to get a new API key")
}

func handleBadRequestError() {
	fmt.Println("💡 Solutions:")
	fmt.Println("   - Check your request parameters")
	fmt.Println("   - Verify required fields are provided")
	fmt.Println("   - Check parameter formats and values")
}

func handleRateLimitError() {
	fmt.Println("💡 Solutions:")
	fmt.Println("   - Wait before making another request")
	fmt.Println("   - Implement exponential backoff")
	fmt.Println("   - Consider upgrading your plan")
}

func handleNotFoundError() {
	fmt.Println("💡 Solutions:")
	fmt.Println("   - Check if the symbol/endpoint exists")
	fmt.Println("   - Verify the URL is correct")
	fmt.Println("   - Check API documentation")
}

func handleTimeoutError() {
	fmt.Println("💡 Solutions:")
	fmt.Println("   - Increase request timeout")
	fmt.Println("   - Check network connectivity")
	fmt.Println("   - Retry the request")
}

func handleNetworkError() {
	fmt.Println("💡 Solutions:")
	fmt.Println("   - Check internet connection")
	fmt.Println("   - Verify firewall settings")
	fmt.Println("   - Check if proxy is required")
}

func handleHTTPError() {
	fmt.Println("💡 Solutions:")
	fmt.Println("   - Check server status")
	fmt.Println("   - Retry after some time")
	fmt.Println("   - Contact API support if persistent")
}

func handleGenericError() {
	fmt.Println("💡 Solutions:")
	fmt.Println("   - Check error details")
	fmt.Println("   - Review API documentation")
	fmt.Println("   - Contact support with error details")
}
