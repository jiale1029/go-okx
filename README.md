# OKX Go SDK (v5)

A powerful, modular, and secure Go SDK for the OKX API v5. Designed for developers building trading bots, portfolio trackers, and financial tools.

## Key Capabilities

### 🛡️ Secure Authentication
- **Automatic HMAC-SHA256 Signing**: All private requests are signed automatically.
- **Environment Variable Support**: Built-in best practices for handling sensitive API keys.
- **Sandbox Mode**: Easy toggle for simulated trading to test your strategies risk-free.

### 📊 REST API Modules
- **Market Data**: Fetch real-time tickers, order books, and candlesticks.
- **Account Management**: Access balances, positions, and leverage configurations.
- **Trade Execution**: Place, cancel, and manage orders (Spot, Swap, Futures, Options).
- **Funding**: Manage deposits, withdrawals, and internal transfers.
- **Sub-accounts**: Complete management for main and sub-account ecosystems.

### 🚀 Future Work (Part 2)
- **WebSocket Client**: Real-time data streaming and private channel subscriptions.

---

## Installation

```bash
go get github.com/jiale1029/go-okx
```

## Quick Start

### 1. Set up your credentials
Create a `.env` file (see [.env.example](.env.example)):
```bash
OKX_API_KEY=your_api_key
OKX_SECRET_KEY=your_secret_key
OKX_PASSPHRASE=your_passphrase
```

### 2. Usage Example

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jiale1029/go-okx/okx"
)

func main() {
	// 1. Initialize the client
	apiKey := os.Getenv("OKX_API_KEY")
	client := okx.NewClient(apiKey, os.Getenv("OKX_SECRET_KEY"), os.Getenv("OKX_PASSPHRASE"), false)

	ctx := context.Background()

	// 2. Fetch Public Market Tickers
	tickers, _ := client.GetTickers(ctx, "SPOT")
	if len(tickers) > 0 {
		fmt.Printf("Symbol: %s, Last Price: %s\n", tickers[0].InstId, tickers[0].Last)
	}

	// 3. Fetch Private Account Balance
	balance, err := client.GetAccountBalance(ctx, "USDT")
	if err == nil && balance != nil {
		fmt.Printf("Total Equity: %s USD\n", balance.TotalEq)
	}
}
```

## Security Best Practices
- **Rotate Keys**: Regularly rotate your OKX API keys.
- **Restrict IP**: In the OKX dashboard, bind your API keys to specific trusted IPs.
- **Never Commit Secrets**: Ensure your `.env` file is in your `.gitignore`.

## License
MIT