package iox

import (
	"io"
)

type (
	textReadOptions struct {
		keepSpace     bool
		withoutBlanks bool
		omitPrefix    string
	}

	TextReadOption func(*textReadOptions)
)

func DupReadCloser(reader io.ReadCloser) (io.ReadCloser, io.ReadCloser) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), *new(io.ReadCloser)
}

func KeepSpace() TextReadOption { _ = "STUB: not implemented"; return *new(TextReadOption) }

func LimitDupReadCloser(reader io.ReadCloser, n int64) (io.ReadCloser, io.ReadCloser) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), *new(io.ReadCloser)
}

func ReadBytes(reader io.Reader, buf []byte) error { _ = "STUB: not implemented"; return nil }

func ReadText(filename string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ReadTextLines(filename string, opts ...TextReadOption) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WithoutBlank() TextReadOption { _ = "STUB: not implemented"; return *new(TextReadOption) }

func OmitWithPrefix(prefix string) TextReadOption {
	_ = "STUB: not implemented"
	return *new(TextReadOption)
}
