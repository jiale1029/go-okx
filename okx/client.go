package okx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jiale1029/go-okx/utils"
)

const (
	DefaultBaseURL = "https://my.okx.com"
)

// Client is the OKX API client
type Client struct {
	BaseURL    string
	APIKey     string
	SecretKey  string
	Passphrase string
	IsSandbox  bool
	HTTPClient *http.Client
}

// NewClient creates a new OKX API client
func NewClient(apiKey, secretKey, passphrase string, isSandbox bool) *Client {
	return &Client{
		BaseURL:    DefaultBaseURL,
		APIKey:     apiKey,
		SecretKey:  secretKey,
		Passphrase: passphrase,
		IsSandbox:  isSandbox,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Request is the internal method to perform HTTP requests
func (c *Client) Do(ctx context.Context, method, path string, body interface{}, private bool) (*http.Response, error) {
	url := c.BaseURL + path
	var bodyBytes []byte
	var err error

	if body != nil {
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.IsSandbox {
		req.Header.Set("x-simulated-trading", "1")
	}

	if private {
		timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
		signature, err := utils.Sign(c.SecretKey, timestamp, method, path, string(bodyBytes))
		if err != nil {
			return nil, fmt.Errorf("failed to sign request: %w", err)
		}

		req.Header.Set("OK-ACCESS-KEY", c.APIKey)
		req.Header.Set("OK-ACCESS-SIGN", signature)
		req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
		req.Header.Set("OK-ACCESS-PASSPHRASE", c.Passphrase)
		fmt.Println(req.Header)
	}

	return c.HTTPClient.Do(req)
}
