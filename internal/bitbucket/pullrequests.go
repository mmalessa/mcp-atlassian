package bitbucket

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"mcp-kit/internal/mcpkit"
)

type PullRequestInput struct {
	Repo string `json:"repo" jsonschema:"Repository slug - the segment after the workspace in the PR URL (e.g. 'my-repo' in bitbucket.org/my-workspace/my-repo/pull-requests/123)"`
	ID   int    `json:"id" jsonschema:"Pull request ID (the number in the PR URL)"`
}

func fetchPR(repo string, id int, subpath string) (*mcp.CallToolResult, any, error) {
	url := prURL(repo, id)
	if subpath != "" {
		url += "/" + subpath
	}
	body, err := get(url)
	if err != nil {
		return nil, nil, err
	}
	return mcpkit.TextResult(string(body)), nil, nil
}

func handleGetPullRequest(_ context.Context, _ *mcp.CallToolRequest, in PullRequestInput) (*mcp.CallToolResult, any, error) {
	return fetchPR(in.Repo, in.ID, "")
}

func handleGetPullRequestDiff(_ context.Context, _ *mcp.CallToolRequest, in PullRequestInput) (*mcp.CallToolResult, any, error) {
	return fetchPR(in.Repo, in.ID, "diff")
}

func handleGetPullRequestDiffstat(_ context.Context, _ *mcp.CallToolRequest, in PullRequestInput) (*mcp.CallToolResult, any, error) {
	return fetchPR(in.Repo, in.ID, "diffstat")
}

func handleGetPullRequestComments(_ context.Context, _ *mcp.CallToolRequest, in PullRequestInput) (*mcp.CallToolResult, any, error) {
	return fetchPR(in.Repo, in.ID, "comments")
}

func handleGetPullRequestCommits(_ context.Context, _ *mcp.CallToolRequest, in PullRequestInput) (*mcp.CallToolResult, any, error) {
	return fetchPR(in.Repo, in.ID, "commits")
}
