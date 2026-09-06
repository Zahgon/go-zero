package ast

import (
	"io"
	"reflect"
)

type FieldFilter func(name string, value reflect.Value) bool

func NotNilFilter(_ string, v reflect.Value) bool { _ = "STUB: not implemented"; return false }

func Fprint(w io.Writer, x interface{}, f FieldFilter) error { _ = "STUB: not implemented"; return nil }

func fprint(w io.Writer, x interface{}, f FieldFilter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func Print(x interface{}) error { _ = "STUB: not implemented"; return nil }

type printer struct {
	output       io.Writer
	filter       FieldFilter
	ptrmap       map[interface{}]int
	prefixIndent int
	last         byte
	line         int
}

var prefixIndent = []byte(".  ")

func (p *printer) Write(data []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

type localError struct {
	err error
}

func (p *printer) printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (p *printer) print(x reflect.Value) { _ = "STUB: not implemented"; return }
