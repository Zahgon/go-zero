package mocksql

import (
	"errors"
	"reflect"
)

const tagName = "db"

var (
	ErrNotMatchDestination = errors.New("not matching destination to scan")

	ErrNotReadableValue = errors.New("value not addressable or interfaceable")

	ErrNotSettable = errors.New("passed in variable is not settable")

	ErrUnsupportedValueType = errors.New("unsupported unmarshal type")
)

type rowsScanner interface {
	Columns() ([]string, error)
	Err() error
	Next() bool
	Scan(v ...any) error
}

func getTaggedFieldValueMap(v reflect.Value) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mapStructFieldsIntoSlice(v reflect.Value, columns []string, strict bool) ([]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseTagName(field reflect.StructField) string { _ = "STUB: not implemented"; return "" }

func unmarshalRow(v any, scanner rowsScanner, strict bool) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalRows(v any, scanner rowsScanner, strict bool) error {
	_ = "STUB: not implemented"
	return nil
}

func unwrapFields(v reflect.Value) []reflect.Value { _ = "STUB: not implemented"; return nil }
