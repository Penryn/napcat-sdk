// Package errorsx 保存 SDK 内部共享的错误类型。
package errorsx

import "errors"

// ErrClosed 表示 client 或底层连接已经关闭。
var ErrClosed = errors.New("napcat: client closed")

// ErrTimeout 表示请求等待响应超时。
var ErrTimeout = errors.New("napcat: request timeout")

// TransportError 表示 HTTP 或 WebSocket 传输层失败。
type TransportError struct {
	Op     string
	Status int
	Body   []byte
	Err    error
}

func (e *TransportError) Error() string {
	if e.Err != nil {
		return "napcat: transport error: " + e.Err.Error()
	}
	return "napcat: transport error"
}

func (e *TransportError) Unwrap() error {
	return e.Err
}

// APIError 表示 NapCat 返回了合法 envelope，但业务状态失败。
type APIError struct {
	Action  string
	Status  string
	RetCode int
	Message string
	Wording string
	Raw     []byte
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return "napcat: api error: " + e.Message
	}
	if e.Wording != "" {
		return "napcat: api error: " + e.Wording
	}
	return "napcat: api error"
}

// ProtocolError 表示 SDK 无法解析 envelope、响应或事件。
type ProtocolError struct {
	Message string
	Raw     []byte
}

func (e *ProtocolError) Error() string {
	if e.Message == "" {
		return "napcat: protocol error"
	}
	return "napcat: protocol error: " + e.Message
}
