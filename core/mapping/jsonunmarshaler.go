package mapping

import (
	"io"
)

const jsonTagKey = "json"

var jsonUnmarshaler = NewUnmarshaler(jsonTagKey)

func UnmarshalJsonBytes(content []byte, v any, opts ...UnmarshalOption) error {
	_ = "STUB: not implemented"
	return nil
}

func UnmarshalJsonMap(m map[string]any, v any, opts ...UnmarshalOption) error {
	_ = "STUB: not implemented"
	return nil
}

func UnmarshalJsonReader(reader io.Reader, v any, opts ...UnmarshalOption) error {
	_ = "STUB: not implemented"
	return nil
}

func getJsonUnmarshaler(opts ...UnmarshalOption) *Unmarshaler {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalJsonBytes(content []byte, v any, unmarshaler *Unmarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalJsonReader(reader io.Reader, v any, unmarshaler *Unmarshaler) error {
	_ = "STUB: not implemented"
	return nil
}
