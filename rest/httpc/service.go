package httpc

import (
	"context"
	"net/http"
)

type (
	Option func(r *http.Request) *http.Request

	Service interface {
		Do(ctx context.Context, method, url string, data any) (*http.Response, error)

		DoRequest(r *http.Request) (*http.Response, error)
	}

	namedService struct {
		name string
		cli  *http.Client
		opts []Option
	}
)

func NewService(name string, opts ...Option) Service {
	_ = "STUB: not implemented"
	return *new(Service)
}

func NewServiceWithClient(name string, cli *http.Client, opts ...Option) Service {
	_ = "STUB: not implemented"
	return *new(Service)
}

func (s namedService) Do(ctx context.Context, method, url string, data any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s namedService) DoRequest(r *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s namedService) do(r *http.Request) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func acceptable(resp *http.Response, err error) bool { _ = "STUB: not implemented"; return false }
