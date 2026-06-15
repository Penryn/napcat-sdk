package transport

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phlin/napcat-sdk/internal/errorsx"
)

func TestHTTPCallerPostsActionWithToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/send_private_msg" {
			t.Fatalf("请求路径 = %s，期望 /send_private_msg", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer token-1" {
			t.Fatalf("Authorization = %q", r.Header.Get("Authorization"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok","retcode":0,"data":{"message_id":42}}`))
	}))
	defer server.Close()

	caller := NewHTTPCaller(server.URL, HTTPOptions{Token: "token-1"})

	var out struct {
		MessageID int64 `json:"message_id"`
	}
	if err := caller.Call(context.Background(), "send_private_msg", map[string]any{"user_id": 1}, &out); err != nil {
		t.Fatalf("HTTP 调用失败: %v", err)
	}
	if out.MessageID != 42 {
		t.Fatalf("message_id = %d，期望 42", out.MessageID)
	}
}

func TestHTTPCallerReturnsAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"failed","retcode":100,"message":"bad request","wording":"参数错误"}`))
	}))
	defer server.Close()

	caller := NewHTTPCaller(server.URL, HTTPOptions{})

	var out struct{}
	err := caller.Call(context.Background(), "get_login_info", nil, &out)
	var apiErr *errorsx.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("错误类型 = %T，期望 APIError", err)
	}
	if apiErr.RetCode != 100 || apiErr.Message != "bad request" {
		t.Fatalf("APIError 字段错误: %#v", apiErr)
	}
}

func TestHTTPCallerReturnsTransportErrorForNon2xx(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusBadGateway)
	}))
	defer server.Close()

	caller := NewHTTPCaller(server.URL, HTTPOptions{})

	var out struct{}
	err := caller.Call(context.Background(), "get_login_info", nil, &out)
	var transportErr *errorsx.TransportError
	if !errors.As(err, &transportErr) {
		t.Fatalf("错误类型 = %T，期望 TransportError", err)
	}
	if transportErr.Status != http.StatusBadGateway {
		t.Fatalf("HTTP 状态码 = %d，期望 502", transportErr.Status)
	}
}
