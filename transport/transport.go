// Package transport 提供 NapCat API 调用的底层传输抽象。
package transport

import "context"

// Caller 是生成 API 层依赖的最小调用接口。
type Caller interface {
	Call(ctx context.Context, action string, params any, result any) error
}
