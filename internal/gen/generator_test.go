package gen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestToExportedNameConvertsActionNames(t *testing.T) {
	tests := map[string]string{
		"send_private_msg":    "SendPrivateMsg",
		"get_login_info":      "GetLoginInfo",
		"ArkSharePeer":        "ArkSharePeer",
		".ocr_image":          "DotOcrImage",
		"_mark_all_as_read":   "UnderscoreMarkAllAsRead",
		"get_group_file_url":  "GetGroupFileURL",
		"get_robot_uin_range": "GetRobotUINRange",
	}

	for input, want := range tests {
		if got := ToExportedName(input); got != want {
			t.Fatalf("ToExportedName(%q) = %q，期望 %q", input, got, want)
		}
	}
}

func TestGenerateContainsRepresentativeActions(t *testing.T) {
	dir := t.TempDir()
	spec := filepath.Join("..", "openapi", "4.18.6", "openapi.json")

	if err := GenerateFromFile(spec, dir); err != nil {
		t.Fatalf("生成 API 失败: %v", err)
	}

	actions := readFile(t, filepath.Join(dir, "actions_gen.go"))
	client := readFile(t, filepath.Join(dir, "client_gen.go"))
	requests := readFile(t, filepath.Join(dir, "requests_gen.go"))
	responses := readFile(t, filepath.Join(dir, "responses_gen.go"))

	for _, want := range []string{
		`ActionSendPrivateMsg Action = "send_private_msg"`,
		`ActionGetLoginInfo Action = "get_login_info"`,
		`ActionGetGroupMemberList Action = "get_group_member_list"`,
	} {
		if !strings.Contains(actions, want) {
			t.Fatalf("actions_gen.go 缺少 %s", want)
		}
	}
	for _, want := range []string{
		"func (c *Client) SendPrivateMsg(",
		"func (c *Client) GetLoginInfo(",
		"func (c *Client) GetGroupMemberList(",
	} {
		if !strings.Contains(client, want) {
			t.Fatalf("client_gen.go 缺少 %s", want)
		}
	}
	if !strings.Contains(requests, "type SendPrivateMsgRequest struct") {
		t.Fatalf("requests_gen.go 缺少 SendPrivateMsgRequest")
	}
	if !strings.Contains(responses, "type GetLoginInfoResponse struct") {
		t.Fatalf("responses_gen.go 缺少 GetLoginInfoResponse")
	}
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("读取文件 %s 失败: %v", path, err)
	}
	return string(data)
}
