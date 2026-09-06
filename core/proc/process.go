package proc

import (
	"os"
	"path/filepath"
)

var (
	procName string
	pid      int
)

func init() {
	procName = filepath.Base(os.Args[0])
	pid = os.Getpid()
}

func Pid() int { _ = "STUB: not implemented"; return 0 }

func ProcessName() string { _ = "STUB: not implemented"; return "" }
