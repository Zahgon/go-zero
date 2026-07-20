package mcp

import (
	"context"

	sdkmcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

type (
	Tool            = sdkmcp.Tool
	CallToolParams  = sdkmcp.CallToolParams
	CallToolResult  = sdkmcp.CallToolResult
	CallToolRequest = sdkmcp.CallToolRequest

	Content      = sdkmcp.Content
	TextContent  = sdkmcp.TextContent
	ImageContent = sdkmcp.ImageContent
	AudioContent = sdkmcp.AudioContent

	Prompt          = sdkmcp.Prompt
	PromptMessage   = sdkmcp.PromptMessage
	GetPromptParams = sdkmcp.GetPromptParams
	GetPromptResult = sdkmcp.GetPromptResult

	Resource           = sdkmcp.Resource
	ResourceContents   = sdkmcp.ResourceContents
	ReadResourceParams = sdkmcp.ReadResourceParams
	ReadResourceResult = sdkmcp.ReadResourceResult

	Server         = sdkmcp.Server
	ServerSession  = sdkmcp.ServerSession
	ServerOptions  = sdkmcp.ServerOptions
	Implementation = sdkmcp.Implementation

	SSEHandler            = sdkmcp.SSEHandler
	StreamableHTTPHandler = sdkmcp.StreamableHTTPHandler
)

type ToolHandler[Args any, Meta any] func(
	ctx context.Context,
	req *CallToolRequest,
	args Args,
) (*CallToolResult, Meta, error)

type PromptHandler func(
	ctx context.Context,
	req *sdkmcp.GetPromptRequest,
	args map[string]string,
) (*GetPromptResult, error)

type ResourceHandler func(
	ctx context.Context,
	req *sdkmcp.ReadResourceRequest,
	uri string,
) (*ReadResourceResult, error)

func AddTool[In, Out any](server McpServer, tool *Tool, handler func(context.Context, *CallToolRequest, In) (*CallToolResult, Out, error)) {
	_ = "STUB: not implemented"
	return
}
