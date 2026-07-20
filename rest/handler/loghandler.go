package handler

import (
	"bufio"
	"bytes"
	"net"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/core/utils"
	"github.com/zeromicro/go-zero/rest/internal"
	"github.com/zeromicro/go-zero/rest/internal/response"
)

const (
	limitBodyBytes          = 1024
	limitDetailedBodyBytes  = 4096
	defaultSlowThreshold    = time.Millisecond * 500
	defaultSSESlowThreshold = time.Minute * 3
)

var (
	slowThreshold    = syncx.ForAtomicDuration(defaultSlowThreshold)
	sseSlowThreshold = syncx.ForAtomicDuration(defaultSSESlowThreshold)
)

func LogHandler(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type detailLoggedResponseWriter struct {
	writer *response.WithCodeResponseWriter
	buf    *bytes.Buffer
}

func newDetailLoggedResponseWriter(writer *response.WithCodeResponseWriter,
	buf *bytes.Buffer) *detailLoggedResponseWriter {
	_ = "STUB: not implemented"
	return nil
}

func (w *detailLoggedResponseWriter) Flush() { _ = "STUB: not implemented"; return }

func (w *detailLoggedResponseWriter) Header() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (w *detailLoggedResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

func (w *detailLoggedResponseWriter) Write(bs []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *detailLoggedResponseWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }

func DetailedLogHandler(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func SetSlowThreshold(threshold time.Duration) { _ = "STUB: not implemented"; return }

func SetSSESlowThreshold(threshold time.Duration) { _ = "STUB: not implemented"; return }

func dumpRequest(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func getSlowThreshold(r *http.Request) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func isOkResponse(code int) bool { _ = "STUB: not implemented"; return false }

func logBrief(r *http.Request, code int, timer *utils.ElapsedTimer, logs *internal.LogCollector) {
	_ = "STUB: not implemented"
	return
}

func logDetails(r *http.Request, response *detailLoggedResponseWriter, timer *utils.ElapsedTimer,
	logs *internal.LogCollector) {
	_ = "STUB: not implemented"
	return
}

func wrapMethod(method string) string { _ = "STUB: not implemented"; return "" }

func wrapStatusCode(code int) string { _ = "STUB: not implemented"; return "" }
