package event

import (
	"testing"
)

func TestParsePrivateMessage(t *testing.T) {
	raw := []byte(`{
		"time": 1718000000,
		"post_type": "message",
		"message_type": "private",
		"sub_type": "friend",
		"self_id": 10000,
		"message_id": 11,
		"user_id": 123456,
		"raw_message": "你好",
		"message": [{"type":"text","data":{"text":"你好"}}],
		"sender": {"user_id":123456,"nickname":"张三","sex":"unknown","age":18}
	}`)

	ev, err := Parse(raw)
	if err != nil {
		t.Fatalf("解析私聊消息失败: %v", err)
	}

	msg, ok := ev.(*PrivateMessage)
	if !ok {
		t.Fatalf("事件类型 = %T，期望 *PrivateMessage", ev)
	}
	if msg.PostType() != "message" || msg.SelfID() != 10000 || msg.UserID != 123456 {
		t.Fatalf("基础字段解析错误: %#v", msg)
	}
	if msg.Message.Text() != "你好" {
		t.Fatalf("消息文本 = %q，期望 你好", msg.Message.Text())
	}
	if msg.Sender.Nickname != "张三" {
		t.Fatalf("发送者昵称 = %q，期望 张三", msg.Sender.Nickname)
	}
}

func TestParseGroupMessage(t *testing.T) {
	raw := []byte(`{
		"time": 1718000001,
		"post_type": "message",
		"message_type": "group",
		"sub_type": "normal",
		"self_id": 10000,
		"message_id": 12,
		"group_id": 654321,
		"user_id": 123456,
		"raw_message": "大家好",
		"message": [{"type":"text","data":{"text":"大家好"}}],
		"sender": {"user_id":123456,"nickname":"李四","card":"群名片","role":"member"}
	}`)

	ev, err := Parse(raw)
	if err != nil {
		t.Fatalf("解析群消息失败: %v", err)
	}

	msg, ok := ev.(*GroupMessage)
	if !ok {
		t.Fatalf("事件类型 = %T，期望 *GroupMessage", ev)
	}
	if msg.GroupID != 654321 || msg.UserID != 123456 {
		t.Fatalf("群消息字段解析错误: %#v", msg)
	}
	if msg.Sender.Card != "群名片" {
		t.Fatalf("群名片 = %q，期望 群名片", msg.Sender.Card)
	}
}

func TestParseUnknownEventPreservesRaw(t *testing.T) {
	raw := []byte(`{"time":1718000002,"post_type":"custom","self_id":10000,"value":1}`)

	ev, err := Parse(raw)
	if err != nil {
		t.Fatalf("解析未知事件失败: %v", err)
	}

	unknown, ok := ev.(*UnknownEvent)
	if !ok {
		t.Fatalf("事件类型 = %T，期望 *UnknownEvent", ev)
	}
	if unknown.PostType() != "custom" || unknown.SelfID() != 10000 {
		t.Fatalf("未知事件基础字段解析错误: %#v", unknown)
	}
	if string(unknown.Raw()) != string(raw) {
		t.Fatalf("未知事件 raw 未保留: %s", unknown.Raw())
	}
}
