package atlassian

import "github.com/modelcontextprotocol/go-sdk/mcp"

func Register(s *mcp.Server) {
	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_jira_task",
		Description: "Get a Jira issue by its key (e.g. PROJ-123).",
	}, handleGetJiraTask)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "search_jira",
		Description: "Search Jira issues using a JQL query.",
	}, handleSearchJira)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_confluence_page",
		Description: "Get a Confluence page by its ID (includes body.storage).",
	}, handleGetConfluencePage)
}
