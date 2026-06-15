// Package jsonx 统一封装 SDK 的 JSON 编解码实现。
package jsonx

import "github.com/bytedance/sonic"

// Marshal 使用 Sonic 编码 JSON。
func Marshal(v any) ([]byte, error) {
	return sonic.Marshal(v)
}

// Unmarshal 使用 Sonic 解码 JSON。
func Unmarshal(data []byte, v any) error {
	return sonic.Unmarshal(data, v)
}
