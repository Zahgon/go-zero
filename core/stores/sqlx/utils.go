package sqlx

import (
	"context"
	"errors"
	"strings"
)

var errUnbalancedEscape = errors.New("no char after escape char")

func desensitize(datasource string) string { _ = "STUB: not implemented"; return "" }

func escape(input string) string { _ = "STUB: not implemented"; return "" }

func format(query string, args ...any) (val string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func logInstanceError(ctx context.Context, datasource string, err error) {
	_ = "STUB: not implemented"
	return
}

func logSqlError(ctx context.Context, stmt string, err error) { _ = "STUB: not implemented"; return }

func writeValue(buf *strings.Builder, arg any) { _ = "STUB: not implemented"; return }

type acceptableError struct {
	err error
}

func newAcceptableError(err error) error { _ = "STUB: not implemented"; return nil }

func (e acceptableError) Error() string { _ = "STUB: not implemented"; return "" }
