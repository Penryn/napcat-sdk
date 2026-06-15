package message

import (
	"testing"

	"github.com/zjutjh/napcat-sdk/internal/jsonx"
)

func TestChainMarshalUnmarshalPreservesSegments(t *testing.T) {
	chain := ChainOf(
		Text("你好"),
		At(123456),
		Image("https://example.com/a.png"),
	)

	data, err := jsonx.Marshal(chain)
	if err != nil {
		t.Fatalf("编码消息链失败: %v", err)
	}

	var got Chain
	if err := jsonx.Unmarshal(data, &got); err != nil {
		t.Fatalf("解码消息链失败: %v", err)
	}

	if len(got) != 3 {
		t.Fatalf("消息段数量 = %d，期望 3", len(got))
	}
	if got[0].Type != "text" || got[1].Type != "at" || got[2].Type != "image" {
		t.Fatalf("消息段类型未保留: %#v", got)
	}
	if got[0].String("text") != "你好" {
		t.Fatalf("文本内容 = %q，期望 你好", got[0].String("text"))
	}
	if got[1].String("qq") != "123456" {
		t.Fatalf("at qq = %q，期望 123456", got[1].String("qq"))
	}
}

func TestChainTextExtractsOnlyTextSegments(t *testing.T) {
	chain := ChainOf(
		Text("ping"),
		At(10001),
		Text(" pong"),
	)

	if got := chain.Text(); got != "ping pong" {
		t.Fatalf("纯文本 = %q，期望 ping pong", got)
	}
}
