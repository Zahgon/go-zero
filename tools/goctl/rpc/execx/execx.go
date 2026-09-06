package execx

import (
	"bytes"
)

type RunFunc func(string, string, ...*bytes.Buffer) (string, error)

func Run(arg, dir string, in ...*bytes.Buffer) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
