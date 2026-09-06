package mcp

import "net/http"

type RequestMetadataExtractor func(*http.Request) RequestMetadata

type McpOption interface {
	apply(*serverOptions)
}

type mcpOptionFunc func(*serverOptions)

func (f mcpOptionFunc) apply(opts *serverOptions) { _ = "STUB: not implemented"; return }

type serverOptions struct {
	requestMetadataExtractor RequestMetadataExtractor
}

func defaultServerOptions() serverOptions { _ = "STUB: not implemented"; return *new(serverOptions) }

func WithRequestMetadataExtractor(extractor RequestMetadataExtractor) McpOption {
	_ = "STUB: not implemented"
	return *new(McpOption)
}
