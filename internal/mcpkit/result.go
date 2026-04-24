package mcpkit

import "github.com/modelcontextprotocol/go-sdk/mcp"

// TextResult wraps text content as an MCP tool result.
func TextResult(text string) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: text}},
	}
}
