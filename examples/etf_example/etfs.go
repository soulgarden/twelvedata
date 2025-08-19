// Package main demonstrates how to retrieve ETF data using the Twelve Data API.
package main

import (
	"fmt"
	"log"

	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata"
	"github.com/soulgarden/twelvedata/request"
	"github.com/valyala/fasthttp"
)

// This is an example function that can be called from another main package
// or run as a separate example. To avoid multiple main() function conflicts,
// this function is named differently.
func runETFExample() {
	logger := zerolog.New(nil).Level(zerolog.Disabled) // Disable logging for example

	cfg := &twelvedata.Conf{
		APIKey: "demo", // Use demo API key instead of hardcoded key
	}

	httpCli := twelvedata.NewHTTPCli(&fasthttp.Client{}, cfg, &logger)
	client := twelvedata.NewClient(httpCli, cfg)

	// Get ETF summary data using the current API
	etfData, credits, err := client.GetETFSummary(request.GetETFSummary{
		APIKey: request.APIKey{
			APIKey: cfg.APIKey,
		},
		Symbol: "VTI", // Example ETF symbol
	})

	if err != nil {
		log.Printf("Error getting ETF data: %v", err)
		return
	}

	fmt.Printf("✅ Successfully retrieved ETF data\n")
	fmt.Printf("📊 Credits used: %d, remaining: %d\n",
		credits.GetCreditsUsed(), credits.GetCreditsLeft())

	// Show ETF information
	fmt.Printf("  - Symbol: %s\n", etfData.ETF.Summary.Symbol)
	fmt.Printf("  - Name: %s\n", etfData.ETF.Summary.Name)
	fmt.Printf("  - Fund Family: %s\n", etfData.ETF.Summary.FundFamily)
	fmt.Printf("  - Currency: %s\n", etfData.ETF.Summary.Currency)
	if etfData.ETF.Summary.NetAssets.Valid {
		fmt.Printf("  - Net Assets: $%.2fB\n", float64(etfData.ETF.Summary.NetAssets.Int64)/1e9)
	}
	if etfData.ETF.Summary.ExpenseRatioNet.Valid {
		fmt.Printf("  - Expense Ratio: %.3f%%\n", etfData.ETF.Summary.ExpenseRatioNet.Float64*100)
	}
}

// This example can be run independently.
func main() { runETFExample() }
