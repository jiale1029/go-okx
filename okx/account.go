package okx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// AccountBalance represents OKX account balance data
type AccountBalance struct {
	Details []BalanceDetail `json:"details"`
	AdjEq   string          `json:"adjEq"`
	Imr     string          `json:"imr"`
	IsoEq   string          `json:"isoEq"`
	MgnRatio string         `json:"mgnRatio"`
	Mmr      string          `json:"mmr"`
	NotionalUsd string      `json:"notionalUsd"`
	OrdFroz string          `json:"ordFroz"`
	TotalEq string          `json:"totalEq"`
	UTime   string          `json:"uTime"`
}

// BalanceDetail represents detailed balance for a specific currency
type BalanceDetail struct {
	AvailBal string `json:"availBal"`
	AvailEq  string `json:"availEq"`
	Ccy      string `json:"ccy"`
	CrossLiab string `json:"crossLiab"`
	DisEq    string `json:"disEq"`
	Eq       string `json:"eq"`
	EqUsd    string `json:"eqUsd"`
	FrozenBal string `json:"frozenBal"`
	Interest string `json:"interest"`
	IsoEq    string `json:"isoEq"`
	IsoLiab  string `json:"isoLiab"`
	IsoUpl   string `json:"isoUpl"`
	Liab     string `json:"liab"`
	MaxLoan  string `json:"maxLoan"`
	MgnRatio string `json:"mgnRatio"`
	NotionalCcy string `json:"notionalCcy"`
	OrdFroz  string `json:"ordFroz"`
	Twap     string `json:"twap"`
	UTime    string `json:"uTime"`
	Upl      string `json:"upl"`
	UplLiab  string `json:"uplLiab"`
	StgyEq   string `json:"stgyEq"`
}

// GetAccountBalance fetches the account balance
// ccy: optional list of currencies
func (c *Client) GetAccountBalance(ctx context.Context, currencies ...string) (*AccountBalance, error) {
	params := url.Values{}
	if len(currencies) > 0 {
		var ccyStr string
		for i, ccy := range currencies {
			if i > 0 {
				ccyStr += ","
			}
			ccyStr += ccy
		}
		params.Add("ccy", ccyStr)
	}

	path := "/api/v5/account/balance"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

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

	if len(base.Data) == 0 {
		return nil, nil
	}

	var balance AccountBalance
	if err := json.Unmarshal(base.Data[0], &balance); err != nil {
		return nil, fmt.Errorf("failed to unmarshal balance: %w", err)
	}

	return &balance, nil
}
