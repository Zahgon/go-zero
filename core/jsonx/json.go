package jsonx

import (
	"encoding/json"
	"io"
)

func Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func MarshalToString(v any) (string, error) { _ = "STUB: not implemented"; return "", nil }

func Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

func UnmarshalFromString(str string, v any) error { _ = "STUB: not implemented"; return nil }

func UnmarshalFromReader(reader io.Reader, v any) error { _ = "STUB: not implemented"; return nil }

func unmarshalUseNumber(decoder *json.Decoder, v any) error { _ = "STUB: not implemented"; return nil }

func formatError(v string, err error) error { _ = "STUB: not implemented"; return nil }
