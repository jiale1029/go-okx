package okx

import (
	"context"
	"encoding/json"
	"fmt"
)

// FundingBalance represents OKX funding account balance data
type FundingBalance struct {
	Ccy      string `json:"ccy"`
	Bal      string `json:"bal"`
	Frozen   string `json:"frozen"`
	AvailBal string `json:"availBal"`
}

// GetFundingBalance fetches the funding account balance
func (c *Client) GetFundingBalance(ctx context.Context, currencies ...string) ([]FundingBalance, error) {
	path := "/api/v5/asset/balances"
	resp, err := c.Do(ctx, "GET", path, nil, true)
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

	var balances []FundingBalance
	for _, item := range base.Data {
		var b FundingBalance
		if err := json.Unmarshal(item, &b); err != nil {
			return nil, fmt.Errorf("failed to unmarshal funding balance: %w", err)
		}
		balances = append(balances, b)
	}

	return balances, nil
}
