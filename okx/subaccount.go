package okx

import (
	"context"
	"encoding/json"
	"fmt"
)

// SubAccount represents OKX sub-account data
type SubAccount struct {
	SubAcct string `json:"subAcct"`
	Status  string `json:"status"`
	Ts      string `json:"ts"`
}

// GetSubAccounts fetches the list of sub-accounts
func (c *Client) GetSubAccounts(ctx context.Context) ([]SubAccount, error) {
	path := "/api/v5/users/subaccount/list"
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

	var subAccounts []SubAccount
	for _, item := range base.Data {
		var s SubAccount
		if err := json.Unmarshal(item, &s); err != nil {
			return nil, fmt.Errorf("failed to unmarshal sub-account: %w", err)
		}
		subAccounts = append(subAccounts, s)
	}

	return subAccounts, nil
}
