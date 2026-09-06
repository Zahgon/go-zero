package httpx

import (
	"net/http"
	"sync"

	"github.com/zeromicro/go-zero/core/mapping"
)

const (
	formKey           = "form"
	pathKey           = "path"
	maxMemory         = 32 << 20
	maxBodyLen        = 8 << 20
	separator         = ";"
	tokensInAttribute = 2
)

var (
	formUnmarshaler = mapping.NewUnmarshaler(
		formKey,
		mapping.WithStringValues(),
		mapping.WithOpaqueKeys(),
		mapping.WithFromArray())
	pathUnmarshaler = mapping.NewUnmarshaler(
		pathKey,
		mapping.WithStringValues(),
		mapping.WithOpaqueKeys())

	validator     Validator
	validatorLock sync.RWMutex
)

type Validator interface {
	Validate(r *http.Request, data any) error
}

func Parse(r *http.Request, v any) error { _ = "STUB: not implemented"; return nil }

func ParseHeaders(r *http.Request, v any) error { _ = "STUB: not implemented"; return nil }

func ParseForm(r *http.Request, v any) error { _ = "STUB: not implemented"; return nil }

func ParseHeader(headerValue string) map[string]string { _ = "STUB: not implemented"; return nil }

func ParseJsonBody(r *http.Request, v any) error { _ = "STUB: not implemented"; return nil }

func ParsePath(r *http.Request, v any) error { _ = "STUB: not implemented"; return nil }

func SetValidator(val Validator) { _ = "STUB: not implemented"; return }

func getValidator() Validator { _ = "STUB: not implemented"; return *new(Validator) }

func withJsonBody(r *http.Request) bool { _ = "STUB: not implemented"; return false }
