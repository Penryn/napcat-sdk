//go:generate go run ../cmd/napcatgen -spec ../internal/openapi/4.18.6/openapi.json -out .

// Package api 提供生成 API 方法共享的运行时类型。
package api

import (
	"context"

	"github.com/phlin/napcat-sdk/transport"
)

// Client 调用 NapCat action。强类型方法由生成代码补充。
type Client struct {
	caller transport.Caller
}

// NewClient 创建 API client。
func NewClient(caller transport.Caller) *Client {
	return &Client{caller: caller}
}

// Call 调用原始 action，供尚未强类型化的新接口兜底。
func (c *Client) Call(ctx context.Context, action string, params any, result any) error {
	return c.caller.Call(ctx, action, params, result)
}
