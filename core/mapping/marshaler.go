package mapping

import (
	"reflect"
)

const (
	emptyTag       = ""
	tagKVSeparator = ":"
)

func Marshal(val any) (map[string]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTag(field reflect.StructField) (string, bool) { _ = "STUB: not implemented"; return "", false }

func insertValue(collector map[string]map[string]any, tag string, key string, val any) {
	_ = "STUB: not implemented"
	return
}

func processMember(field reflect.StructField, value reflect.Value,
	collector map[string]map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func validate(field reflect.StructField, value reflect.Value, opt *fieldOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOptional(field reflect.StructField, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOptions(value reflect.Value, opt *fieldOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRange(value reflect.Value, opt *fieldOptions) error {
	_ = "STUB: not implemented"
	return nil
}
