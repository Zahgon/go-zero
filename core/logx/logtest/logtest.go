package logtest

import (
	"bytes"
	"testing"
)

type Buffer struct {
	buf *bytes.Buffer
	t   *testing.T
}

func Discard(t *testing.T) { _ = "STUB: not implemented"; return }

func NewCollector(t *testing.T) *Buffer { _ = "STUB: not implemented"; return nil }

func (b *Buffer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *Buffer) Content() string { _ = "STUB: not implemented"; return "" }

func (b *Buffer) Reset() { _ = "STUB: not implemented"; return }

func (b *Buffer) String() string { _ = "STUB: not implemented"; return "" }

func PanicOnFatal(t *testing.T) { _ = "STUB: not implemented"; return }
