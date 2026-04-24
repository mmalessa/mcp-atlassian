package main

import (
	"log"

	"mcp-kit/internal/atlassian"
	"mcp-kit/internal/mcpkit"
)

func main() {
	app := mcpkit.App{
		Name:     "mcp-atlassian",
		Version:  "0.1.0",
		Register: atlassian.Register,
	}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
