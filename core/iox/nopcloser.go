package iox

import "io"

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error { _ = "STUB: not implemented"; return nil }

func NopCloser(w io.Writer) io.WriteCloser { _ = "STUB: not implemented"; return *new(io.WriteCloser) }
