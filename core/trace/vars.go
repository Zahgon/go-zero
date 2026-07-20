package trace

import "net/http"

var TraceIdKey = http.CanonicalHeaderKey("x-trace-id")
