//go:build linux || darwin || freebsd

package proc

import (
	"os"
)

const (
	goroutineProfile = "goroutine"
	debugLevel       = 2
)

type creator interface {
	Create(name string) (file *os.File, err error)
}

func dumpGoroutines(ctor creator) { _ = "STUB: not implemented"; return }

type fileCreator struct{}

func (fc fileCreator) Create(name string) (file *os.File, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
