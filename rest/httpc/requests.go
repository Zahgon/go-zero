package httpc

import (
	"context"
	"net/http"
	nurl "net/url"

	"github.com/zeromicro/go-zero/rest/httpc/internal"
)

var interceptors = []internal.Interceptor{
	internal.LogInterceptor,
}

func Do(ctx context.Context, method, url string, data any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DoRequest(r *http.Request) (*http.Response, error) { _ = "STUB: not implemented"; return nil, nil }

type (
	client interface {
		do(r *http.Request) (*http.Response, error)
	}

	defaultClient struct{}
)

func (c defaultClient) do(r *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildFormQuery(u *nurl.URL, val map[string]any) string { _ = "STUB: not implemented"; return "" }

func buildRequest(ctx context.Context, method, url string, data any) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fillHeader(r *http.Request, val map[string]any) { _ = "STUB: not implemented"; return }

func fillPath(u *nurl.URL, val map[string]any) error { _ = "STUB: not implemented"; return nil }

func request(r *http.Request, cli client) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
