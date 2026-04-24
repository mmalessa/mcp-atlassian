package main

import (
	"log"

	"mcp-kit/internal/bitbucket"
	"mcp-kit/internal/mcpkit"
)

func main() {
	app := mcpkit.App{
		Name:     "mcp-bitbucket",
		Version:  "0.1.0",
		Register: bitbucket.Register,
	}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
