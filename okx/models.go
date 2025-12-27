package okx

import (
	"encoding/json"
)

// BaseResponse is the common structure for OKX API responses
type BaseResponse struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data []json.RawMessage `json:"data"`
}

// ErrorResponse represents an OKX API error
type ErrorResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (e *ErrorResponse) Error() string {
	return e.Msg
}
