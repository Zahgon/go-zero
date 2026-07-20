package sysx

import "go.uber.org/automaxprocs/maxprocs"

func init() {
	maxprocs.Set(maxprocs.Logger(nil))
}
