package okx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// Ticker represents OKX ticker data
type Ticker struct {
	InstType  string `json:"instType"`
	InstId    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24h   string `json:"open24h"`
	High24h   string `json:"high24h"`
	Low24h    string `json:"low24h"`
	VolCcy24h string `json:"volCcy24h"`
	Vol24h    string `json:"vol24h"`
	Ts        string `json:"ts"`
	SodPx     string `json:"sodPx"`
}

// GetTickers fetches tickers for a specific instrument type
// instType: SPOT, MARGIN, SWAP, FUTURES, OPTION
func (c *Client) GetTickers(ctx context.Context, instType string) ([]Ticker, error) {
	params := url.Values{}
	params.Add("instType", instType)

	path := "/api/v5/market/tickers?" + params.Encode()
	resp, err := c.Do(ctx, "GET", path, nil, false)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var base BaseResponse
	if err := json.NewDecoder(resp.Body).Decode(&base); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if base.Code != "0" {
		return nil, &ErrorResponse{Code: base.Code, Msg: base.Msg}
	}

	var tickers []Ticker
	for _, item := range base.Data {
		var t Ticker
		if err := json.Unmarshal(item, &t); err != nil {
			return nil, fmt.Errorf("failed to unmarshal ticker: %w", err)
		}
		tickers = append(tickers, t)
	}

	return tickers, nil
}
