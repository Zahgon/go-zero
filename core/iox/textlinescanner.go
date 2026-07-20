package iox

import (
	"bufio"
	"io"
)

type TextLineScanner struct {
	reader  *bufio.Reader
	hasNext bool
	line    string
	err     error
}

func NewTextLineScanner(reader io.Reader) *TextLineScanner { _ = "STUB: not implemented"; return nil }

func (scanner *TextLineScanner) Scan() bool { _ = "STUB: not implemented"; return false }

func (scanner *TextLineScanner) Line() (string, error) { _ = "STUB: not implemented"; return "", nil }
