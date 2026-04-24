package atlassian

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"mcp-kit/internal/mcpkit"
)

type GetJiraTaskInput struct {
	IssueKey string `json:"issue_key" jsonschema:"Jira issue key, e.g. PROJ-123"`
}

type SearchJiraInput struct {
	JQL string `json:"jql" jsonschema:"JQL query string, e.g. \"assignee = currentUser() ORDER BY created DESC\""`
}

func handleGetJiraTask(_ context.Context, _ *mcp.CallToolRequest, in GetJiraTaskInput) (*mcp.CallToolResult, any, error) {
	url := fmt.Sprintf("%s/rest/api/3/issue/%s",
		os.Getenv("ATLASSIAN_BASE_URL"),
		in.IssueKey,
	)

	body, err := atlassianRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	return mcpkit.TextResult(string(body)), nil, nil
}

func handleSearchJira(_ context.Context, _ *mcp.CallToolRequest, in SearchJiraInput) (*mcp.CallToolResult, any, error) {
	url := fmt.Sprintf("%s/rest/api/3/search/jql", os.Getenv("ATLASSIAN_BASE_URL"))

	reqBody, err := json.Marshal(map[string]any{
		"jql":    in.JQL,
		"fields": []string{"*all"},
	})
	if err != nil {
		return nil, nil, err
	}

	body, err := atlassianRequest("POST", url, reqBody)
	if err != nil {
		return nil, nil, err
	}

	return mcpkit.TextResult(string(body)), nil, nil
}
