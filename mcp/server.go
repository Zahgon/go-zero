package mcp

import (
	"net/http"

	sdkmcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/zeromicro/go-zero/rest"
)

type McpServer interface {
	Start()

	Stop()
}

type mcpServerImpl struct {
	conf       McpConf
	httpServer *rest.Server
	mcpServer  *sdkmcp.Server
	options    serverOptions
}

func NewMcpServer(c McpConf) McpServer { _ = "STUB: not implemented"; return *new(McpServer) }

func NewMcpServerWithOptions(c McpConf, opts ...McpOption) McpServer {
	_ = "STUB: not implemented"
	return *new(McpServer)
}

func (s *mcpServerImpl) Start() { _ = "STUB: not implemented"; return }

func (s *mcpServerImpl) Stop() { _ = "STUB: not implemented"; return }

func (s *mcpServerImpl) setupSSETransport() { _ = "STUB: not implemented"; return }

func (s *mcpServerImpl) setupStreamableTransport() { _ = "STUB: not implemented"; return }

func (s *mcpServerImpl) registerRoutes(handler http.Handler, endpoint string) {
	_ = "STUB: not implemented"
	return
}

func (s *mcpServerImpl) wrapRequestMetadata(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
