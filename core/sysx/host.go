package sysx

import (
	"os"

	"github.com/zeromicro/go-zero/core/stringx"
)

var hostname string

func init() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		hostname = stringx.RandId()
	}
}

func Hostname() string { _ = "STUB: not implemented"; return "" }
