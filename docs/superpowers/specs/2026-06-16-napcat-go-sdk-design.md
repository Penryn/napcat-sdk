# NapCat Go SDK 设计规格

日期：2026-06-16

## 目标定位

构建一个纯 Go SDK，用于对接 NapCat/OneBot 11。SDK 的核心目标是强类型、OpenAPI 自动生成、轻量运行时，以及稳定的事件解析体验。

这个 SDK 只负责：

- 调用 NapCat 4.18.6 API。
- 通过 HTTP 或 WebSocket 与 NapCat 通信。
- 解析 OneBot 事件。
- 构造和解析消息段。

这个 SDK 不做机器人框架，不内置插件系统、命令系统、路由系统或业务运行时。后续如果需要框架，可以基于本 SDK 单独构建。

## 资料来源

- NapCat 官方 API 版本页：<https://napneko.github.io/api/4.18.6>
- NapCat OneBot 网络通信文档：<https://napneko.github.io/onebot/network>
- NapCat OneBot API 文档：<https://napneko.github.io/onebot/api>
- NapCat 扩展 API 文档：<https://napneko.github.io/develop/api/doc>
- 官方 OpenAPI 来源：`NapCatDocs/src/api/4.18.6/openapi.json`
- 参考仓库：
  - <https://github.com/nekoite/go-napcat>
  - <https://github.com/nsxdevx/nsxbot>
  - <https://github.com/ncatbot/NcatBot>
  - <https://github.com/faithleysath/napcat-sdk>

官方 OpenAPI 文件是 API action 的主来源。本次调研确认 `4.18.6/openapi.json` 使用 OpenAPI 3.1.0，标题为 `NapCat OneBot 11 HTTP API`，包含 170 个 path 和 39 个 schema。

## 设计目标

1. 为 NapCat 4.18.6 API 提供强类型 Go 方法。
2. 从官方 OpenAPI 生成 API 请求类型、响应类型和 action 方法。
3. HTTP 调用使用 Resty。
4. WebSocket 支持正向连接和反向 WebSocket server。
5. OneBot 上报事件解析为明确的 Go 类型。
6. 提供易用的消息段构造器和解析器。
7. 运行时依赖保持克制。
8. 生成代码可通过 `go generate` 复现。

## 非目标

1. 不做插件系统。
2. 不做命令解析器。
3. 不做中间件或路由框架。
4. 不做定时任务、持久化、管理后台、CLI bot runtime 或 NapCat 进程管理。
5. 不绑定 `zap`、`logrus` 等具体日志库。
6. 普通用户 API 不暴露 Resty 或 WebSocket 底层实现细节。

## 技术栈

最低 Go 版本：`go 1.24`。

运行时依赖：

- HTTP：`github.com/go-resty/resty/v2`
- WebSocket：`github.com/coder/websocket`
- JSON：`github.com/bytedance/sonic`
- 日志：标准库 `log/slog`，可选配置，默认静默

生成期依赖：

- OpenAPI 解析：`github.com/getkin/kin-openapi`

测试依赖：

- 标准库 `testing`、`httptest`、`net/http`、`context`
- 只有在断言和 mock 明显更清晰时才引入 `github.com/stretchr/testify`

JSON 必须统一经过一个很薄的内部封装，避免运行时代码到处直接调用 `sonic`：

```go
package jsonx

import "github.com/bytedance/sonic"

func Marshal(v any) ([]byte, error) {
	return sonic.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return sonic.Unmarshal(data, v)
}
```

HTTP transport 内部使用 Resty，但请求体和响应体的 JSON 编解码统一走 `jsonx`。Resty 负责 base URL、header、超时、重试、HTTP 状态码和请求发送。

## 包结构

建议项目结构：

```text
napcat-sdk/
  api/                  OpenAPI 生成的请求/响应结构体和 API 方法
  event/                手写 OneBot 事件类型和解析器
  message/              手写消息段、消息链、CQ 辅助
  transport/            transport 接口和高级构造器
  internal/jsonx/       sonic 封装
  internal/openapi/     固化的官方 OpenAPI spec
  internal/gen/         生成器内部实现，不进入 runtime
  cmd/napcatgen/        OpenAPI 生成器命令
  examples/
```

根包 `napcat` 是用户主入口：

```go
client, err := napcat.DialWebSocket(ctx, "ws://127.0.0.1:3001", napcat.WithToken(token))
if err != nil {
	return err
}
defer client.Close()

resp, err := client.API().SendPrivateMsg(ctx, api.SendPrivateMsgRequest{
	UserID:  123456,
	Message: message.Text("pong"),
})
```

## Client 模型

`Client` 持有一个 transport，并暴露强类型 API、raw call、生命周期和事件流。

```go
type Client struct {
	api    *api.Client
	events <-chan event.Event
}

func NewHTTPClient(baseURL string, opts ...Option) (*Client, error)
func DialWebSocket(ctx context.Context, url string, opts ...Option) (*Client, error)
func ServeReverseWebSocket(ctx context.Context, addr string, handler func(*Client), opts ...Option) error

func (c *Client) API() *api.Client
func (c *Client) Events() <-chan event.Event
func (c *Client) Call(ctx context.Context, action string, params any, result any) error
func (c *Client) Close() error
```

`Client` 必须支持并发 API 调用。事件消费由应用负责 fan-out，SDK 不为每个事件无限制创建 goroutine。

## 配置项

Option 同时服务 HTTP 和 WebSocket transport。

```go
func WithToken(token string) Option
func WithHTTPTimeout(timeout time.Duration) Option
func WithHTTPRetry(count int, wait time.Duration) Option
func WithRestyClient(client *resty.Client) Option
func WithWebSocketDialOptions(opts *websocket.DialOptions) Option
func WithRequestTimeout(timeout time.Duration) Option
func WithEventBuffer(size int) Option
func WithLogger(logger *slog.Logger) Option
```

`WithRestyClient` 只面向高级用户。普通用户不需要导入 Resty。

## Transport 抽象

生成的 `api.Client` 面向一个内部调用接口：

```go
type Caller interface {
	Call(ctx context.Context, action string, params any, result any) error
}
```

HTTP transport 将 `action` 映射为 `POST /{action}`，请求体为 `params` 的 JSON。

典型 OneBot 响应 envelope：

```json
{
  "status": "ok",
  "retcode": 0,
  "data": {},
  "message": "",
  "wording": ""
}
```

WebSocket transport 发送：

```json
{
  "action": "send_group_msg",
  "params": {},
  "echo": "sdk-generated-id"
}
```

收到 WebSocket frame 时：

- 有 `echo`：按 `echo` 路由到等待中的 API 调用。
- 无 `echo`：按 OneBot 上报事件解析，送入 `Client.Events()`。

## HTTP Transport

HTTP transport 内部使用 `resty.Client`。

要求：

- 配置 token 时设置 `Authorization: Bearer <token>`。
- 请求地址为 `baseURL + "/" + action`。
- 请求体通过 `jsonx.Marshal` 编码。
- 响应体通过 `jsonx.Unmarshal` 解码。
- 非 2xx HTTP 状态码返回 transport error。
- OneBot envelope 中 `status != "ok"` 或 `retcode != 0` 返回 API error。
- 错误中保留原始响应 body，方便排查。
- 支持 `context.Context` 取消和超时。

普通 SDK API 不返回 Resty 的 response 对象。

## WebSocket Transport

正向 WebSocket 模式连接 NapCat 提供的 WebSocket 地址，例如 `ws://127.0.0.1:3001`。

要求：

- 配置 token 时设置 `Authorization: Bearer <token>`。
- 维护按 `echo` 索引的 pending request map。
- 生成不冲突的 echo ID。
- 支持并发调用。
- 收到响应、超时、context 取消或连接关闭时清理 pending request。
- API 响应 frame 和事件 frame 分流处理。
- 连接关闭时，所有 pending 调用返回 `ErrClosed`。
- 事件 channel 必须有边界。默认策略是阻塞 reader，不静默丢事件。

反向 WebSocket server 模式由 SDK 监听地址，等待 NapCat 主动连接。每个连接包装成一个具有同样 API 和事件流能力的 `Client`。

## API 生成策略

SDK 使用自研生成器，而不是直接套通用 OpenAPI client generator。原因是 NapCat 的同一套 action 需要同时支持 HTTP 和 WebSocket，两者共享的是 `action + params -> envelope(data)` 调用模型。

生成命令：

```sh
go generate ./...
```

生成器相关路径：

```text
cmd/napcatgen/
internal/gen/
internal/openapi/4.18.6/openapi.json
```

生成文件：

```text
api/actions_gen.go
api/client_gen.go
api/requests_gen.go
api/responses_gen.go
api/types_gen.go
```

生成方法形态：

```go
func (c *Client) SendPrivateMsg(ctx context.Context, req SendPrivateMsgRequest) (*SendPrivateMsgResponse, error) {
	var out SendPrivateMsgResponse
	if err := c.caller.Call(ctx, ActionSendPrivateMsg, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
```

生成代码要求：

- 保留 JSON 字段名和 `omitempty` 语义。
- 将 `/send_private_msg` 转换为 `ActionSendPrivateMsg`。
- 尽可能把 OpenAPI summary 转为 Go doc comment。
- 文件头写明 NapCat OpenAPI 版本和生成命令。
- 不覆盖手写文件。
- 输出确定性稳定，CI 可以通过 diff 判断生成产物是否过期。

OpenAPI 中可能存在宽松或不一致 schema。生成器应优先生成有用的强类型结构体；遇到无法可靠表达的字段时，可以退回 `[]byte` 或 SDK 内部的 `RawMessage` alias。

## Event 包

事件类型手写，因为事件由运行时字段判别，且 Go 用户需要自然的 type switch。

核心接口：

```go
type Event interface {
	PostType() string
	SelfID() int64
	Time() int64
	Raw() []byte
}
```

解析入口：

```go
func Parse(data []byte) (Event, error)
```

首版覆盖：

- `PrivateMessage`
- `GroupMessage`
- `Notice`
- `Request`
- `Meta`
- `MessageSent`
- `UnknownEvent`

更细的 notice/request/meta 类型可以逐步增加，但基础 parser 必须先保证未知事件不丢失。未知字段保留在 raw payload 中，便于兼容新版 NapCat。

消息事件结构示例：

```go
type PrivateMessage struct {
	Base
	MessageID  int64
	UserID     int64
	Message    message.Chain
	RawMessage string
	Sender     PrivateSender
}

type GroupMessage struct {
	Base
	MessageID  int64
	GroupID    int64
	UserID     int64
	Message    message.Chain
	RawMessage string
	Sender     GroupSender
}
```

## Message 包

消息包手写，目标是 Go 调用体验清晰。

核心类型：

```go
type Chain []Segment

type Segment struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}
```

构造器示例：

```go
message.Text("hello")
message.At(123456)
message.AtAll()
message.Image("https://example.com/a.png")
message.Reply(10001)
message.Record("file://voice.silk")
message.Video("https://example.com/a.mp4")
message.Face("14")
message.NodeCustom("nickname", 123456, message.Text("nested"))
```

`Chain` 支持：

- JSON array marshal/unmarshal。
- 提取纯文本。
- 按类型过滤 segment。
- 可行范围内支持 CQ 字符串解析/格式化。
- 安全保留未知 segment 类型。

消息构造器的返回类型必须保持一致。API request 中接收消息的字段要能编码 NapCat 支持的 string、segment 或 chain。

## 错误处理

定义 sentinel error 和结构化 error：

```go
var ErrClosed = errors.New("napcat: client closed")
var ErrTimeout = errors.New("napcat: request timeout")

type TransportError struct {
	Op     string
	Status int
	Body   []byte
	Err    error
}

type APIError struct {
	Action  string
	Status  string
	RetCode int
	Message string
	Wording string
	Raw     []byte
}

type ProtocolError struct {
	Message string
	Raw     []byte
}
```

规则：

- Transport error 表示 HTTP/WebSocket 传输失败。
- API error 表示 NapCat 返回了合法 envelope，但业务状态失败。
- Protocol error 表示 SDK 无法解析 envelope、响应或事件。
- context cancellation 和 deadline 尽量原样返回 `context.Canceled` 或 `context.DeadlineExceeded`。

## 版本策略

模块从 `v0` 开始。

首版生成 API 固定到 NapCat `4.18.6`。后续支持新版本有两种策略：

1. 替换固定 OpenAPI spec，并为下一版 SDK 重新生成。
2. 只有在确实需要同时支持多个 NapCat API 版本时，才引入 versioned generated package。

默认选择第一种：SDK release line 只维护一套生成 API，并在 README 中明确对应的 NapCat OpenAPI 版本。

## 测试策略

单元测试：

- JSON codec wrapper。
- 消息段 marshal/unmarshal。
- 事件 parser 判别和未知事件兜底。
- 错误分类。
- 生成的 action 名称和方法映射。

HTTP 测试：

- 使用 `httptest.Server` 校验请求 path、authorization header、body 和 context 取消。
- 成功 envelope 能解码为强类型 response。
- 失败 envelope 转为 `APIError`。
- 非 2xx 响应转为 `TransportError`。

WebSocket 测试：

- 本地测试 server 接受连接并校验 action frame。
- echo response 能唤醒对应 pending call。
- 无 echo 的 event frame 进入 `Events()`。
- 并发调用不会串响应。
- 关闭连接时 pending call 返回 `ErrClosed`。

生成器测试：

- 加载官方 OpenAPI 4.18.6。
- 校验生成 action 数量符合预期。
- 校验核心、群组、用户、文件、系统、扩展、流式接口均有代表性方法。
- 对小型 fixture spec 做 golden-file 测试。
- CI 执行 `go generate ./...`，生成文件有差异则失败。

## 文档和示例

必须提供以下示例：

1. HTTP API 调用：
   - `NewHTTPClient`
   - `GetLoginInfo`
   - `SendPrivateMsg`

2. 正向 WebSocket 事件循环：
   - `DialWebSocket`
   - 消费 `Events()`
   - 对 `/ping` 回复 `pong`

3. 反向 WebSocket server：
   - 接收 NapCat 连接
   - 消费事件
   - 通过同一个 client 调 API

4. 富媒体消息：
   - 文本
   - at
   - 图片
   - 回复

5. Raw escape hatch：
   - `client.Call(ctx, "new_action", params, &result)`

## 验收标准

实现完成应满足：

1. `go generate ./...` 可以从 NapCat 4.18.6 OpenAPI 生成确定性 API 代码。
2. 生成 API 为官方 OpenAPI 中所有 path 暴露强类型方法。
3. HTTP client 通过 Resty 调用 NapCat action，并使用 Sonic 做 JSON 编解码。
4. WebSocket client 支持并发调用、echo 匹配和事件流。
5. 反向 WebSocket server 将每个连接包装成 `Client`。
6. 事件 parser 能处理已知消息事件，并保留未知事件。
7. Message 包能构造和解码常见 OneBot 消息段。
8. 测试覆盖 HTTP、WebSocket、事件解析、消息 JSON、错误处理和生成器确定性。
9. README 包含最小 HTTP 示例和最小 WebSocket 事件循环示例。

## 已确定决策

- SDK 形态：纯 SDK，不做机器人框架。
- API 覆盖：OpenAPI 生成优先。
- HTTP：Resty。
- WebSocket：coder/websocket。
- JSON：Sonic。
- OpenAPI：使用 `kin-openapi` 自研生成器，不直接生成通用 OpenAPI client。
