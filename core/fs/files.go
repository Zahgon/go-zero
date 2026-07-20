//go:build linux || darwin || freebsd

package fs

import (
	"os"
)

func CloseOnExec(file *os.File) { _ = "STUB: not implemented"; return }
