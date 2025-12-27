package okx

import (
	"context"
	"encoding/json"
	"fmt"
)

// OrderRequest represents a request to place an order
type OrderRequest struct {
	InstId     string `json:"instId"`
	TdMode     string `json:"tdMode"`
	Side       string `json:"side"`
	OrdType    string `json:"ordType"`
	Sz         string `json:"sz"`
	Px         string `json:"px,omitempty"`
	ClOrdId    string `json:"clOrdId,omitempty"`
	Tag        string `json:"tag,omitempty"`
	PosSide    string `json:"posSide,omitempty"`
	ReduceOnly bool   `json:"reduceOnly,omitempty"`
}

// OrderResponse represents the response from placing an order
type OrderResponse struct {
	OrdId   string `json:"ordId"`
	ClOrdId string `json:"clOrdId"`
	Tag     string `json:"tag"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}

// PlaceOrder places an order
func (c *Client) PlaceOrder(ctx context.Context, req *OrderRequest) (*OrderResponse, error) {
	path := "/api/v5/trade/order"
	resp, err := c.Do(ctx, "POST", path, req, true)
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

	var orderResp OrderResponse
	if err := json.Unmarshal(base.Data[0], &orderResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal order response: %w", err)
	}

	return &orderResp, nil
}
