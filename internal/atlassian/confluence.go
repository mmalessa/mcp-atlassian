package atlassian

import (
	"context"
	"fmt"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"mcp-kit/internal/mcpkit"
)

type GetConfluencePageInput struct {
	PageID string `json:"page_id" jsonschema:"Confluence page ID"`
}

func handleGetConfluencePage(_ context.Context, _ *mcp.CallToolRequest, in GetConfluencePageInput) (*mcp.CallToolResult, any, error) {
	url := fmt.Sprintf("%s/wiki/rest/api/content/%s?expand=body.storage",
		os.Getenv("ATLASSIAN_BASE_URL"),
		in.PageID,
	)

	body, err := atlassianRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	return mcpkit.TextResult(string(body)), nil, nil
}
