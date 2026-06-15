// Package gen 实现 NapCat OpenAPI 到 Go API 绑定的生成器。
package gen

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/zjutjh/napcat-sdk/internal/jsonx"
)

// GenerateFromFile 从 OpenAPI spec 生成 API 代码。
func GenerateFromFile(specPath string, outDir string) error {
	loader := openapi3.NewLoader()
	if _, err := loader.LoadFromFile(specPath); err != nil {
		return fmt.Errorf("加载 OpenAPI spec 失败: %w", err)
	}

	data, err := os.ReadFile(specPath)
	if err != nil {
		return fmt.Errorf("读取 OpenAPI spec 失败: %w", err)
	}
	var doc document
	if err := jsonx.Unmarshal(data, &doc); err != nil {
		return fmt.Errorf("解析 OpenAPI spec 失败: %w", err)
	}

	actions := collectActions(doc)
	if len(actions) == 0 {
		return fmt.Errorf("OpenAPI spec 中没有可生成的 action")
	}

	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("创建输出目录失败: %w", err)
	}

	files := map[string][]byte{
		"actions_gen.go":   generateActions(actions),
		"client_gen.go":    generateClient(actions),
		"requests_gen.go":  generateRequests(actions),
		"responses_gen.go": generateResponses(actions),
		"types_gen.go":     generateTypes(),
	}
	for name, content := range files {
		formatted, err := format.Source(content)
		if err != nil {
			return fmt.Errorf("格式化 %s 失败: %w\n%s", name, err, content)
		}
		if err := os.WriteFile(filepath.Join(outDir, name), formatted, 0o644); err != nil {
			return fmt.Errorf("写入 %s 失败: %w", name, err)
		}
	}
	return nil
}

type document struct {
	Paths      map[string]pathItem `json:"paths"`
	Components components          `json:"components"`
}

type components struct {
	Schemas map[string]*schema `json:"schemas"`
}

type pathItem struct {
	Post *operation `json:"post"`
}

type operation struct {
	Summary     string              `json:"summary"`
	RequestBody *requestBody        `json:"requestBody"`
	Responses   map[string]response `json:"responses"`
}

type requestBody struct {
	Content map[string]mediaType `json:"content"`
}

type response struct {
	Content map[string]mediaType `json:"content"`
}

type mediaType struct {
	Schema *schema `json:"schema"`
}

type schema struct {
	Ref         string             `json:"$ref"`
	Type        any                `json:"type"`
	Description string             `json:"description"`
	Properties  map[string]*schema `json:"properties"`
	Items       *schema            `json:"items"`
	Required    []string           `json:"required"`
	AnyOf       []*schema          `json:"anyOf"`
	AllOf       []*schema          `json:"allOf"`
}

type actionSpec struct {
	Action         string
	Name           string
	Summary        string
	RequestFields  []fieldSpec
	ResponseFields []fieldSpec
}

type fieldSpec struct {
	Name        string
	JSONName    string
	Type        string
	Description string
	Required    bool
}

func collectActions(doc document) []actionSpec {
	paths := make([]string, 0, len(doc.Paths))
	for path := range doc.Paths {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	actions := make([]actionSpec, 0, len(paths))
	usedNames := map[string]int{}
	for _, path := range paths {
		item := doc.Paths[path]
		if item.Post == nil {
			continue
		}
		action := strings.TrimPrefix(path, "/")
		name := uniqueName(ToExportedName(action), usedNames)
		actions = append(actions, actionSpec{
			Action:         action,
			Name:           name,
			Summary:        item.Post.Summary,
			RequestFields:  fieldsFromSchema(requestSchema(item.Post), doc.Components.Schemas),
			ResponseFields: fieldsFromSchema(responseDataSchema(item.Post), doc.Components.Schemas),
		})
	}
	return actions
}

func requestSchema(op *operation) *schema {
	if op == nil || op.RequestBody == nil {
		return nil
	}
	if mt, ok := op.RequestBody.Content["application/json"]; ok {
		return mt.Schema
	}
	return nil
}

func responseDataSchema(op *operation) *schema {
	if op == nil {
		return nil
	}
	resp, ok := op.Responses["200"]
	if !ok {
		resp, ok = op.Responses["default"]
		if !ok {
			return nil
		}
	}
	mt, ok := resp.Content["application/json"]
	if !ok || mt.Schema == nil {
		return nil
	}
	return findDataSchema(mt.Schema)
}

func findDataSchema(s *schema) *schema {
	if s == nil {
		return nil
	}
	if data := s.Properties["data"]; data != nil {
		return data
	}
	for _, child := range s.AllOf {
		if data := findDataSchema(child); data != nil {
			return data
		}
	}
	return nil
}

func fieldsFromSchema(s *schema, components map[string]*schema) []fieldSpec {
	s = resolveSchema(s, components)
	if s == nil || len(s.Properties) == 0 {
		return nil
	}
	required := map[string]bool{}
	for _, name := range s.Required {
		required[name] = true
	}

	keys := make([]string, 0, len(s.Properties))
	for key := range s.Properties {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	fields := make([]fieldSpec, 0, len(keys))
	used := map[string]int{}
	for _, key := range keys {
		prop := s.Properties[key]
		name := uniqueName(fieldName(key), used)
		fields = append(fields, fieldSpec{
			Name:        name,
			JSONName:    key,
			Type:        goType(prop, components),
			Description: sanitizeComment(prop.Description),
			Required:    required[key],
		})
	}
	return fields
}

func resolveSchema(s *schema, components map[string]*schema) *schema {
	if s == nil || s.Ref == "" {
		return s
	}
	name := refName(s.Ref)
	if components == nil {
		return s
	}
	if resolved := components[name]; resolved != nil {
		return resolved
	}
	return s
}

func goType(s *schema, components map[string]*schema) string {
	s = resolveSchema(s, components)
	if s == nil {
		return "any"
	}
	if len(s.AnyOf) > 0 {
		return "any"
	}
	switch schemaType(s) {
	case "string":
		return "string"
	case "integer", "number":
		return "int64"
	case "boolean":
		return "bool"
	case "array":
		if s.Items == nil {
			return "[]any"
		}
		itemType := goType(s.Items, components)
		if itemType == "map[string]any" {
			return "[]map[string]any"
		}
		if itemType == "any" {
			return "[]any"
		}
		return "[]" + itemType
	case "object":
		return "map[string]any"
	default:
		return "any"
	}
}

func schemaType(s *schema) string {
	switch value := s.Type.(type) {
	case string:
		return value
	case []any:
		for _, item := range value {
			if text, ok := item.(string); ok && text != "null" {
				return text
			}
		}
	}
	if len(s.Properties) > 0 {
		return "object"
	}
	return ""
}

func generateActions(actions []actionSpec) []byte {
	var b bytes.Buffer
	writeHeader(&b)
	b.WriteString("package api\n\n")
	b.WriteString("// Action 是 NapCat API action 名称。\n")
	b.WriteString("type Action string\n\n")
	b.WriteString("const (\n")
	for _, action := range actions {
		writeComment(&b, "Action"+action.Name, action.Summary)
		fmt.Fprintf(&b, "\tAction%s Action = %q\n", action.Name, action.Action)
	}
	b.WriteString(")\n")
	return b.Bytes()
}

func generateRequests(actions []actionSpec) []byte {
	var b bytes.Buffer
	writeHeader(&b)
	b.WriteString("package api\n\n")
	for _, action := range actions {
		writeComment(&b, action.Name+"Request", action.Summary+" 请求参数。")
		fmt.Fprintf(&b, "type %sRequest struct {\n", action.Name)
		writeFields(&b, action.RequestFields)
		b.WriteString("}\n\n")
	}
	return b.Bytes()
}

func generateResponses(actions []actionSpec) []byte {
	var b bytes.Buffer
	writeHeader(&b)
	b.WriteString("package api\n\n")
	for _, action := range actions {
		writeComment(&b, action.Name+"Response", action.Summary+" 响应数据。")
		fmt.Fprintf(&b, "type %sResponse struct {\n", action.Name)
		writeFields(&b, action.ResponseFields)
		b.WriteString("}\n\n")
	}
	return b.Bytes()
}

func generateClient(actions []actionSpec) []byte {
	var b bytes.Buffer
	writeHeader(&b)
	b.WriteString("package api\n\n")
	b.WriteString("import \"context\"\n\n")
	for _, action := range actions {
		writeComment(&b, action.Name, action.Summary)
		fmt.Fprintf(&b, "func (c *Client) %s(ctx context.Context, req %sRequest) (*%sResponse, error) {\n", action.Name, action.Name, action.Name)
		fmt.Fprintf(&b, "\tvar out %sResponse\n", action.Name)
		fmt.Fprintf(&b, "\tif err := c.caller.Call(ctx, string(Action%s), req, &out); err != nil {\n", action.Name)
		b.WriteString("\t\treturn nil, err\n")
		b.WriteString("\t}\n")
		b.WriteString("\treturn &out, nil\n")
		b.WriteString("}\n\n")
	}
	return b.Bytes()
}

func generateTypes() []byte {
	var b bytes.Buffer
	writeHeader(&b)
	b.WriteString("package api\n\n")
	b.WriteString("// RawMap 表示 OpenAPI 中无法稳定展开的对象字段。\n")
	b.WriteString("type RawMap = map[string]any\n")
	return b.Bytes()
}

func writeFields(b *bytes.Buffer, fields []fieldSpec) {
	for _, field := range fields {
		if field.Description != "" {
			fmt.Fprintf(b, "\t// %s %s\n", field.Name, field.Description)
		}
		tag := field.JSONName
		if !field.Required {
			tag += ",omitempty"
		}
		fmt.Fprintf(b, "\t%s %s `json:%q`\n", field.Name, field.Type, tag)
	}
}

func writeHeader(b *bytes.Buffer) {
	b.WriteString("// 代码由 napcatgen 根据 NapCat OpenAPI 4.18.6 生成；请勿手动修改。\n\n")
}

func writeComment(b *bytes.Buffer, name string, text string) {
	text = sanitizeComment(text)
	if text == "" {
		fmt.Fprintf(b, "// %s 由 NapCat OpenAPI 生成。\n", name)
		return
	}
	fmt.Fprintf(b, "// %s %s\n", name, text)
}

func sanitizeComment(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\r", " ")
	return text
}

func uniqueName(name string, used map[string]int) string {
	if name == "" {
		name = "Value"
	}
	used[name]++
	if used[name] == 1 {
		return name
	}
	return fmt.Sprintf("%s%d", name, used[name])
}

// ToExportedName 将 action 名称转换为导出的 Go 名称。
func ToExportedName(action string) string {
	prefix := ""
	if strings.HasPrefix(action, ".") {
		prefix = "Dot"
	}
	if strings.HasPrefix(action, "_") {
		prefix = "Underscore"
	}
	tokens := splitName(action)
	var b strings.Builder
	b.WriteString(prefix)
	for _, token := range tokens {
		b.WriteString(normalizeToken(token))
	}
	if b.Len() == 0 {
		return "Action"
	}
	return b.String()
}

func fieldName(name string) string {
	out := ToExportedName(name)
	if out == "ID" {
		return "ID"
	}
	return out
}

var splitRegexp = regexp.MustCompile(`[^A-Za-z0-9]+`)

func splitName(name string) []string {
	raw := splitRegexp.Split(name, -1)
	out := make([]string, 0, len(raw))
	for _, token := range raw {
		if token != "" {
			out = append(out, token)
		}
	}
	return out
}

func normalizeToken(token string) string {
	upper := strings.ToUpper(token)
	switch upper {
	case "ID", "URL", "API", "JSON", "QQ", "UIN", "UID", "QID", "IP":
		return upper
	}
	runes := []rune(token)
	if len(runes) == 0 {
		return ""
	}
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func refName(ref string) string {
	idx := strings.LastIndex(ref, "/")
	if idx == -1 {
		return ref
	}
	return ref[idx+1:]
}
