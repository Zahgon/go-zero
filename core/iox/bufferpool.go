package iox

import (
	"bytes"
	"sync"
)

type BufferPool struct {
	capability int
	pool       *sync.Pool
}

func NewBufferPool(capability int) *BufferPool { _ = "STUB: not implemented"; return nil }

func (bp *BufferPool) Get() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

func (bp *BufferPool) Put(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }
