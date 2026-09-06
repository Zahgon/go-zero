package mon

import (
	"context"
	"time"
)

const mongoAddrSep = ","

func FormatAddr(hosts []string) string { _ = "STUB: not implemented"; return "" }

func logDuration(ctx context.Context, name, method string, startTime time.Duration, err error) {
	_ = "STUB: not implemented"
	return
}

func logDurationWithDocs(ctx context.Context, name, method string, startTime time.Duration,
	err error, docs ...any) {
	_ = "STUB: not implemented"
	return
}
