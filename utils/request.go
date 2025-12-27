package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// Sign generates the OKX API signature
// SecretKey: the API secret
// Timestamp: ISO 8601 timestamp
// Method: HTTP method (GET, POST, etc.)
// RequestPath: the relative path of the request
// Body: the request body string
func Sign(secretKey, timestamp, method, requestPath, body string) (string, error) {
	message := timestamp + method + requestPath + body
	h := hmac.New(sha256.New, []byte(secretKey))
	_, err := h.Write([]byte(message))
	if err != nil {
		return "", fmt.Errorf("failed to write hmac: %w", err)
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
