# NapCat Go SDK 实现计划

> **给自动化执行者：** 实现本计划时需要使用 `superpowers:subagent-driven-development`（推荐）或 `superpowers:executing-plans`。步骤使用 checkbox（`- [ ]`）记录执行状态。

**目标：** 实现 `docs/superpowers/specs/2026-06-16-napcat-go-sdk-design.md` 中定义的 NapCat Go SDK。

**架构：** SDK 由手写运行时和 OpenAPI 生成层组成。运行时负责 JSON、错误类型、消息段、事件解析、HTTP/WS transport 和根包 `napcat.Client`；生成层负责把 NapCat OpenAPI action 转为强类型 Go 方法，并通过统一的 `Caller` 接口调用底层 transport。

**技术栈：** Go 1.24、Resty、coder/websocket、Sonic、kin-openapi、标准库 `testing`/`httptest`。

---

## 文件结构

- 创建 `go.mod`：模块信息和依赖。
- 创建 `.gitignore`：忽略本地、系统和构建产物。
- 创建 `internal/jsonx/jsonx.go`：Sonic 封装。
- 创建 `errors.go`：公开错误类型。
- 创建 `message/message.go` 和 `message/message_test.go`：消息段和消息链。
- 创建 `event/event.go` 和 `event/event_test.go`：事件解析和消息事件类型。
- 创建 `transport/transport.go`：共享 `Caller` 接口。
- 创建 `transport/http.go` 和 `transport/http_test.go`：基于 Resty 的 HTTP 调用器。
- 创建 `transport/ws.go` 和 `transport/ws_test.go`：基于 coder/websocket 的 WebSocket 调用器和事件流。
- 创建 `client.go`：根包构造器和 client facade。
- 创建 `api/client.go`：生成 API 的运行时壳。
- 创建 `cmd/napcatgen/main.go` 和 `internal/gen/generator.go`：OpenAPI 生成器。
- 创建 `internal/openapi/4.18.6/openapi.json`：固定版本的官方 NapCat OpenAPI spec。
- 生成 `api/actions_gen.go`、`api/client_gen.go`、`api/requests_gen.go`、`api/responses_gen.go`、`api/types_gen.go`。
- 创建 `README.md` 和 `examples/*`：中文使用文档和示例。

## 任务 1：仓库和模块初始化

**文件：**
- 创建：`go.mod`
- 创建：`.gitignore`
- 修改：`docs/superpowers/plans/2026-06-16-napcat-go-sdk-implementation.md`

- [ ] **步骤 1：创建模块信息**

```go
module github.com/zjutjh/napcat-sdk

go 1.24
```

- [ ] **步骤 2：添加忽略规则**

```gitignore
.DS_Store
bin/
coverage.out
*.test
```

- [ ] **步骤 3：整理模块**

运行：`go mod tidy`

期望：命令退出码为 `0`。在尚未创建 Go 包时，`go.mod` 可暂时不记录依赖。

- [ ] **步骤 4：提交初始化**

运行：`git add go.mod .gitignore docs/superpowers/plans/2026-06-16-napcat-go-sdk-implementation.md && git commit -m "chore: initialize Go module"`

## 任务 2：JSON、错误、消息和事件运行时

**文件：**
- 创建：`internal/jsonx/jsonx.go`
- 创建：`errors.go`
- 创建：`message/message.go`
- 创建：`message/message_test.go`
- 创建：`event/event.go`
- 创建：`event/event_test.go`

- [ ] **步骤 1：先写消息和事件测试**

测试覆盖：

```go
func TestChainMarshalUnmarshalPreservesSegments(t *testing.T)
func TestChainTextExtractsOnlyTextSegments(t *testing.T)
func TestParsePrivateMessage(t *testing.T)
func TestParseGroupMessage(t *testing.T)
func TestParseUnknownEventPreservesRaw(t *testing.T)
```

- [ ] **步骤 2：运行红灯测试**

运行：`go test ./message ./event`

期望：失败，原因是包或符号尚未实现。

- [ ] **步骤 3：实现最小运行时代码**

实现：

```go
package jsonx

func Marshal(v any) ([]byte, error)
func Unmarshal(data []byte, v any) error
```

```go
package message

type Segment struct { Type string; Data any }
type Chain []Segment
func Text(text string) Segment
func At(qq int64) Segment
func AtAll() Segment
func Image(file string) Segment
func Reply(id int64) Segment
func ChainOf(segments ...Segment) Chain
func (c Chain) Text() string
```

```go
package event

type Event interface { PostType() string; SelfID() int64; Time() int64; Raw() []byte }
func Parse(data []byte) (Event, error)
type PrivateMessage struct { Base; MessageID int64; UserID int64; Message message.Chain; RawMessage string; Sender PrivateSender }
type GroupMessage struct { Base; MessageID int64; GroupID int64; UserID int64; Message message.Chain; RawMessage string; Sender GroupSender }
type UnknownEvent struct { Base }
```

- [ ] **步骤 4：运行绿灯测试**

运行：`go test ./message ./event`

期望：通过。

- [ ] **步骤 5：提交运行时基础**

运行：`git add internal/jsonx errors.go message event && git commit -m "feat: add message and event runtime"`

## 任务 3：Transport 和根 Client

**文件：**
- 创建：`transport/transport.go`
- 创建：`transport/http.go`
- 创建：`transport/http_test.go`
- 创建：`transport/ws.go`
- 创建：`transport/ws_test.go`
- 创建：`client.go`

- [ ] **步骤 1：先写 HTTP transport 测试**

测试覆盖：

```go
func TestHTTPCallerPostsActionWithToken(t *testing.T)
func TestHTTPCallerReturnsAPIError(t *testing.T)
func TestHTTPCallerReturnsTransportErrorForNon2xx(t *testing.T)
```

- [ ] **步骤 2：先写 WebSocket transport 测试**

测试覆盖：

```go
func TestWebSocketCallerMatchesEchoResponse(t *testing.T)
func TestWebSocketCallerPublishesEventFrame(t *testing.T)
```

- [ ] **步骤 3：运行红灯测试**

运行：`go test ./transport`

期望：失败，原因是 transport 尚未实现。

- [ ] **步骤 4：实现 transport 和 client**

实现：

```go
package transport

type Caller interface {
	Call(ctx context.Context, action string, params any, result any) error
}
```

HTTP 调用器要求：使用 Resty、通过 Sonic 编解码 JSON、解析 OneBot envelope、支持 token header、非 2xx 错误和 API 错误。

WebSocket 调用器要求：使用 coder/websocket、维护 echo map、提供事件 channel、支持并发调用、支持 context 取消和关闭处理。

根包实现：

```go
func NewHTTPClient(baseURL string, opts ...Option) (*Client, error)
func DialWebSocket(ctx context.Context, url string, opts ...Option) (*Client, error)
func (c *Client) API() *api.Client
func (c *Client) Events() <-chan event.Event
func (c *Client) Call(ctx context.Context, action string, params any, result any) error
func (c *Client) Close() error
```

- [ ] **步骤 5：运行绿灯测试**

运行：`go test ./transport ./...`

期望：transport 通过；如 API/generator 包尚未创建，只允许出现与尚未实现阶段直接相关的失败。

- [ ] **步骤 6：提交 transport**

运行：`git add transport client.go && git commit -m "feat: add HTTP and WebSocket transports"`

## 任务 4：OpenAPI 生成器和生成 API

**文件：**
- 创建：`internal/openapi/4.18.6/openapi.json`
- 创建：`api/client.go`
- 创建：`cmd/napcatgen/main.go`
- 创建：`internal/gen/generator.go`
- 创建：`internal/gen/generator_test.go`
- 生成：`api/actions_gen.go`
- 生成：`api/client_gen.go`
- 生成：`api/requests_gen.go`
- 生成：`api/responses_gen.go`
- 生成：`api/types_gen.go`

- [ ] **步骤 1：复制官方 OpenAPI spec**

运行：`mkdir -p internal/openapi/4.18.6 && cp /tmp/napcat-sdk-reference-repos/NapCatDocs/src/api/4.18.6/openapi.json internal/openapi/4.18.6/openapi.json`

期望：文件存在，且内容为 OpenAPI 3.1.0。

- [ ] **步骤 2：先写生成器测试**

测试覆盖：

```go
func TestGenerateContainsRepresentativeActions(t *testing.T)
func TestToExportedNameConvertsActionNames(t *testing.T)
```

- [ ] **步骤 3：运行红灯测试**

运行：`go test ./internal/gen`

期望：失败，原因是生成器尚未实现。

- [ ] **步骤 4：实现生成器**

生成器行为：

- 读取 `internal/openapi/4.18.6/openapi.json`。
- 枚举所有 paths。
- 将 action name 转换为导出的 Go 名称。
- 生成 action 常量。
- 在能可靠表达时，从 JSON object properties 生成 request struct。
- 对松散 response schema 使用可用的强类型或 `any` 回退。
- 生成基于 `transport.Caller` 的 client 方法。
- 对 paths 和 schema keys 排序，保证输出确定性。

- [ ] **步骤 5：生成 API**

运行：`go run ./cmd/napcatgen -spec internal/openapi/4.18.6/openapi.json -out api`

期望：生成文件存在，并包含 `ActionSendPrivateMsg`、`SendPrivateMsg`、`ActionGetLoginInfo`、`GetLoginInfo`。

- [ ] **步骤 6：运行 API 测试**

运行：`go test ./internal/gen ./api`

期望：通过。

- [ ] **步骤 7：提交生成器和生成 API**

运行：`git add internal/openapi api cmd internal/gen && git commit -m "feat: generate NapCat API bindings"`

## 任务 5：中文文档、示例和完整验证

**文件：**
- 创建：`README.md`
- 创建：`examples/http/main.go`
- 创建：`examples/ws/main.go`
- 创建：`examples/reverse-ws/main.go`

- [ ] **步骤 1：添加中文 README 示例**

README 必须展示：

```go
client, err := napcat.NewHTTPClient("http://127.0.0.1:3000", napcat.WithToken(token))
login, err := client.API().GetLoginInfo(ctx, api.GetLoginInfoRequest{})
```

以及：

```go
client, err := napcat.DialWebSocket(ctx, "ws://127.0.0.1:3001", napcat.WithToken(token))
for ev := range client.Events() {
	switch e := ev.(type) {
	case *event.PrivateMessage:
		_, _ = client.API().SendPrivateMsg(ctx, api.SendPrivateMsgRequest{
			UserID: e.UserID,
			Message: message.Text("pong"),
		})
	}
}
```

- [ ] **步骤 2：添加可运行示例**

在 `examples/` 下添加最小 HTTP、正向 WebSocket、反向 WebSocket 示例。文档和注释使用中文。

- [ ] **步骤 3：运行完整验证**

运行：`go test ./...`

期望：通过。

运行：`go generate ./...`

期望：生成文件不变化。

运行：`git diff --exit-code`

期望：生成后没有 diff。

- [ ] **步骤 4：提交文档和示例**

运行：`git add README.md examples && git commit -m "docs: add SDK usage examples"`

