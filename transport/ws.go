package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/coder/websocket"
	"github.com/phlin/napcat-sdk/event"
	"github.com/phlin/napcat-sdk/internal/errorsx"
	"github.com/phlin/napcat-sdk/internal/jsonx"
)

// WebSocketOptions 配置 WebSocket 调用器。
type WebSocketOptions struct {
	Token          string
	RequestTimeout time.Duration
	EventBuffer    int
	DialOptions    *websocket.DialOptions
}

// WebSocketCaller 使用 OneBot WebSocket action 模型调用 NapCat。
type WebSocketCaller struct {
	conn           *websocket.Conn
	requestTimeout time.Duration
	events         chan event.Event
	pending        map[string]chan envelope
	pendingMu      sync.Mutex
	closed         chan struct{}
	closeOnce      sync.Once
	seq            atomic.Int64
}

type wsRequest struct {
	Action string `json:"action"`
	Params any    `json:"params"`
	Echo   string `json:"echo"`
}

type echoProbe struct {
	Echo string `json:"echo"`
}

// DialWebSocket 连接正向 WebSocket，并启动后台读循环。
func DialWebSocket(ctx context.Context, url string, opts WebSocketOptions) (*WebSocketCaller, error) {
	dialOpts := opts.DialOptions
	if dialOpts == nil {
		dialOpts = &websocket.DialOptions{}
	}
	if opts.Token != "" {
		if dialOpts.HTTPHeader == nil {
			dialOpts.HTTPHeader = http.Header{}
		}
		dialOpts.HTTPHeader.Set("Authorization", "Bearer "+opts.Token)
	}
	conn, _, err := websocket.Dial(ctx, url, dialOpts)
	if err != nil {
		return nil, &errorsx.TransportError{Op: "websocket dial", Err: err}
	}
	return NewWebSocketCaller(conn, opts), nil
}

// NewWebSocketCaller 从已有连接创建调用器，供反向 WebSocket server 复用。
func NewWebSocketCaller(conn *websocket.Conn, opts WebSocketOptions) *WebSocketCaller {
	eventBuffer := opts.EventBuffer
	if eventBuffer <= 0 {
		eventBuffer = 16
	}
	timeout := opts.RequestTimeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	c := &WebSocketCaller{
		conn:           conn,
		requestTimeout: timeout,
		events:         make(chan event.Event, eventBuffer),
		pending:        make(map[string]chan envelope),
		closed:         make(chan struct{}),
	}
	go c.readLoop()
	return c
}

// Events 返回事件流。
func (c *WebSocketCaller) Events() <-chan event.Event {
	return c.events
}

// Call 通过 WebSocket action 调用 NapCat。
func (c *WebSocketCaller) Call(ctx context.Context, action string, params any, result any) error {
	echo := strconv.FormatInt(c.seq.Add(1), 10)
	respCh := make(chan envelope, 1)
	if err := c.addPending(echo, respCh); err != nil {
		return err
	}
	defer c.removePending(echo)

	payload, err := jsonx.Marshal(wsRequest{Action: action, Params: params, Echo: echo})
	if err != nil {
		return fmt.Errorf("编码 WebSocket 请求失败: %w", err)
	}
	if err := c.conn.Write(ctx, websocket.MessageText, payload); err != nil {
		return &errorsx.TransportError{Op: action, Err: err}
	}

	waitCtx := ctx
	cancel := func() {}
	if _, ok := ctx.Deadline(); !ok && c.requestTimeout > 0 {
		waitCtx, cancel = context.WithTimeout(ctx, c.requestTimeout)
	}
	defer cancel()

	select {
	case env := <-respCh:
		raw, _ := jsonx.Marshal(env)
		return decodeEnvelope(action, raw, result)
	case <-waitCtx.Done():
		if errors.Is(waitCtx.Err(), context.DeadlineExceeded) && ctx.Err() == nil {
			return errorsx.ErrTimeout
		}
		return waitCtx.Err()
	case <-c.closed:
		return errorsx.ErrClosed
	}
}

// Close 关闭连接并唤醒所有 pending 调用。
func (c *WebSocketCaller) Close() error {
	var err error
	c.closeOnce.Do(func() {
		close(c.closed)
		err = c.conn.Close(websocket.StatusNormalClosure, "")
		c.pendingMu.Lock()
		for echo, ch := range c.pending {
			delete(c.pending, echo)
			close(ch)
		}
		c.pendingMu.Unlock()
		close(c.events)
	})
	return err
}

func (c *WebSocketCaller) addPending(echo string, ch chan envelope) error {
	select {
	case <-c.closed:
		return errorsx.ErrClosed
	default:
	}
	c.pendingMu.Lock()
	defer c.pendingMu.Unlock()
	c.pending[echo] = ch
	return nil
}

func (c *WebSocketCaller) removePending(echo string) {
	c.pendingMu.Lock()
	delete(c.pending, echo)
	c.pendingMu.Unlock()
}

func (c *WebSocketCaller) readLoop() {
	defer c.Close()
	for {
		_, data, err := c.conn.Read(context.Background())
		if err != nil {
			return
		}
		var probe echoProbe
		if err := jsonx.Unmarshal(data, &probe); err == nil && probe.Echo != "" {
			var env envelope
			if err := jsonx.Unmarshal(data, &env); err != nil {
				continue
			}
			c.pendingMu.Lock()
			ch := c.pending[probe.Echo]
			c.pendingMu.Unlock()
			if ch != nil {
				ch <- env
			}
			continue
		}
		ev, err := event.Parse(data)
		if err != nil {
			continue
		}
		select {
		case c.events <- ev:
		case <-c.closed:
			return
		}
	}
}

func (e envelope) MarshalJSON() ([]byte, error) {
	type alias envelope
	if e.Data == nil {
		return json.Marshal(alias(e))
	}
	return json.Marshal(alias(e))
}
