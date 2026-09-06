package httpc

import (
	"net/http"
)

func Parse(resp *http.Response, val any) error { _ = "STUB: not implemented"; return nil }

func ParseHeaders(resp *http.Response, val any) error { _ = "STUB: not implemented"; return nil }

func ParseJsonBody(resp *http.Response, val any) error { _ = "STUB: not implemented"; return nil }

func isContentTypeJson(r *http.Response) bool { _ = "STUB: not implemented"; return false }
