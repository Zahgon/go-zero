package conf

import (
	"reflect"

	"github.com/zeromicro/go-zero/core/mapping"
)

const (
	jsonTagKey = "json"
	jsonTagSep = ','
)

var (
	fillDefaultUnmarshaler = mapping.NewUnmarshaler(jsonTagKey, mapping.WithDefault())
	loaders                = map[string]func([]byte, any) error{
		".json":  LoadFromJsonBytes,
		".json5": LoadFromJson5Bytes,
		".toml":  LoadFromTomlBytes,
		".yaml":  LoadFromYamlBytes,
		".yml":   LoadFromYamlBytes,
	}
)

type fieldInfo struct {
	children map[string]*fieldInfo
	mapField *fieldInfo
}

func FillDefault(v any) error { _ = "STUB: not implemented"; return nil }

func Load(file string, v any, opts ...Option) error { _ = "STUB: not implemented"; return nil }

func LoadConfig(file string, v any, opts ...Option) error { _ = "STUB: not implemented"; return nil }

func LoadFromJsonBytes(content []byte, v any) error { _ = "STUB: not implemented"; return nil }

func LoadConfigFromJsonBytes(content []byte, v any) error { _ = "STUB: not implemented"; return nil }

func LoadFromTomlBytes(content []byte, v any) error { _ = "STUB: not implemented"; return nil }

func LoadFromYamlBytes(content []byte, v any) error { _ = "STUB: not implemented"; return nil }

func LoadFromJson5Bytes(content []byte, v any) error { _ = "STUB: not implemented"; return nil }

func LoadConfigFromYamlBytes(content []byte, v any) error { _ = "STUB: not implemented"; return nil }

func MustLoad(path string, v any, opts ...Option) { _ = "STUB: not implemented"; return }

func addOrMergeFields(info *fieldInfo, key string, child *fieldInfo, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func buildAnonymousFieldInfo(info *fieldInfo, lowerCaseName string, ft reflect.Type, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func buildFieldsInfo(tp reflect.Type, fullName string) (*fieldInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildNamedFieldInfo(info *fieldInfo, lowerCaseName string, ft reflect.Type, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func buildStructFieldsInfo(tp reflect.Type, fullName string) (*fieldInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTagName(field reflect.StructField) string { _ = "STUB: not implemented"; return "" }

func mergeFields(prev *fieldInfo, children map[string]*fieldInfo, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func toLowerCase(s string) string { _ = "STUB: not implemented"; return "" }

func toLowerCaseInterface(v any, info *fieldInfo) any { _ = "STUB: not implemented"; return *new(any) }

func toLowerCaseKeyMap(m map[string]any, info *fieldInfo) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

type conflictKeyError struct {
	key string
}

func newConflictKeyError(key string) conflictKeyError {
	_ = "STUB: not implemented"
	return *new(conflictKeyError)
}

func (e conflictKeyError) Error() string { _ = "STUB: not implemented"; return "" }

func getFullName(parent, child string) string { _ = "STUB: not implemented"; return "" }
