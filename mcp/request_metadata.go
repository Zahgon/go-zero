package mcp

import (
	"context"
	"net/http"
)

type RequestMetadata struct {
	Headers map[string][]string
	Query   map[string][]string
	Path    map[string]string
}

type requestMetadataCtxKey struct{}

func RequestMetadataFromContext(ctx context.Context) (RequestMetadata, bool) {
	_ = "STUB: not implemented"
	return *new(RequestMetadata), false
}

func HeaderFromContext(ctx context.Context, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func QueryFromContext(ctx context.Context, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func PathFromContext(ctx context.Context, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func requestMetadataFromContext(ctx context.Context) (RequestMetadata, bool) {
	_ = "STUB: not implemented"
	return *new(RequestMetadata), false
}

func DefaultRequestMetadataExtractor(r *http.Request) RequestMetadata {
	_ = "STUB: not implemented"
	return *new(RequestMetadata)
}

func normalizeRequestMetadata(metadata RequestMetadata) RequestMetadata {
	_ = "STUB: not implemented"
	return *new(RequestMetadata)
}

func cloneHeaderValues(values map[string][]string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func cloneCanonicalHeaderValues(values map[string][]string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func clonePathVars(values map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
