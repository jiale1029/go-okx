package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jiale1029/go-okx/okx"
)

func main() {
	// Initialize client using environment variables for safety
	apiKey := os.Getenv("OKX_API_KEY")
	secretKey := os.Getenv("OKX_SECRET_KEY")
	passphrase := os.Getenv("OKX_PASSPHRASE")

	if apiKey == "" || secretKey == "" || passphrase == "" {
		log.Fatal("Missing required environment variables: OKX_API_KEY, OKX_SECRET_KEY, OKX_PASSPHRASE")
	}

	client := okx.NewClient(apiKey, secretKey, passphrase, false)

	ctx := context.Background()

	// 1. Fetch Market Tickers (Public)
	fmt.Println("--- Market Tickers ---")
	tickers, err := client.GetTickers(ctx, "SPOT")
	if err != nil {
		log.Fatalf("Error fetching tickers: %v", err)
	}

	if len(tickers) > 0 {
		t := tickers[0]
		fmt.Printf("Symbol: %s, Last Price: %s, 24h Vol: %s\n", t.InstId, t.Last, t.Vol24h)
	}

	// 2. Fetch Account Balance (Private - requires valid API keys)
	fmt.Println("\n--- Account Balance ---")
	balance, err := client.GetAccountBalance(ctx, "USDT", "BTC")
	if err != nil {
		fmt.Printf("Error fetching balance (likely invalid keys): %v\n", err)
	} else if balance != nil {
		fmt.Printf("Total Equity (USD): %s\n", balance.TotalEq)
		for _, detail := range balance.Details {
			fmt.Printf("Asset: %s, Available: %s\n", detail.Ccy, detail.AvailBal)
		}
	}

	// 3. Place a Limit Order (Private - commented out for safety)
	/*
		orderReq := &okx.OrderRequest{
			InstId:  "BTC-USDT-SWAP",
			TdMode:  "cross",
			Side:    "buy",
			OrdType: "limit",
			Sz:      "1",
			Px:      "50000",
		}
		orderResp, err := client.PlaceOrder(ctx, orderReq)
		if err != nil {
			log.Printf("Error placing order: %v\n", err)
		} else {
			fmt.Printf("Order Placed! ID: %s\n", orderResp.OrdId)
		}
	*/
}
