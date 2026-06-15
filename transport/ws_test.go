package transport

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/zjutjh/napcat-sdk/event"
	"github.com/zjutjh/napcat-sdk/internal/jsonx"
)

func TestWebSocketCallerMatchesEchoResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			t.Fatalf("接受 WebSocket 失败: %v", err)
		}
		defer conn.Close(websocket.StatusNormalClosure, "")

		_, data, err := conn.Read(r.Context())
		if err != nil {
			t.Fatalf("读取请求失败: %v", err)
		}
		var req struct {
			Action string         `json:"action"`
			Params map[string]any `json:"params"`
			Echo   string         `json:"echo"`
		}
		if err := jsonx.Unmarshal(data, &req); err != nil {
			t.Fatalf("解析请求失败: %v", err)
		}
		if req.Action != "get_login_info" || req.Echo == "" {
			t.Fatalf("请求内容错误: %#v", req)
		}
		resp := `{"status":"ok","retcode":0,"data":{"user_id":10000},"echo":"` + req.Echo + `"}`
		if err := conn.Write(r.Context(), websocket.MessageText, []byte(resp)); err != nil {
			t.Fatalf("写入响应失败: %v", err)
		}
	}))
	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	caller, err := DialWebSocket(ctx, "ws"+strings.TrimPrefix(server.URL, "http"), WebSocketOptions{})
	if err != nil {
		t.Fatalf("连接 WebSocket 失败: %v", err)
	}
	defer caller.Close()

	var out struct {
		UserID int64 `json:"user_id"`
	}
	if err := caller.Call(ctx, "get_login_info", map[string]any{}, &out); err != nil {
		t.Fatalf("WebSocket 调用失败: %v", err)
	}
	if out.UserID != 10000 {
		t.Fatalf("user_id = %d，期望 10000", out.UserID)
	}
}

func TestWebSocketCallerPublishesEventFrame(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			t.Fatalf("接受 WebSocket 失败: %v", err)
		}
		defer conn.Close(websocket.StatusNormalClosure, "")

		eventJSON := []byte(`{"time":1718000000,"post_type":"message","message_type":"private","self_id":10000,"message_id":1,"user_id":123,"raw_message":"ping","message":[{"type":"text","data":{"text":"ping"}}],"sender":{"user_id":123,"nickname":"张三"}}`)
		if err := conn.Write(r.Context(), websocket.MessageText, eventJSON); err != nil {
			t.Fatalf("写入事件失败: %v", err)
		}
		<-r.Context().Done()
	}))
	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	caller, err := DialWebSocket(ctx, "ws"+strings.TrimPrefix(server.URL, "http"), WebSocketOptions{EventBuffer: 1})
	if err != nil {
		t.Fatalf("连接 WebSocket 失败: %v", err)
	}
	defer caller.Close()

	select {
	case ev := <-caller.Events():
		if _, ok := ev.(*event.PrivateMessage); !ok {
			t.Fatalf("事件类型 = %T，期望 *event.PrivateMessage", ev)
		}
	case <-ctx.Done():
		t.Fatalf("等待事件超时: %v", ctx.Err())
	}
}
