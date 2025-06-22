package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/douglarek/gist-mcp/prompt"
	"github.com/douglarek/gist-mcp/resource"
	"github.com/google/go-github/v72/github"
	"github.com/mark3labs/mcp-go/server"
)

var (
	debugMode = flag.Bool("debug", false, "debug mode")
	slogLevel = new(slog.LevelVar)
)

func init() {
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slogLevel})
	slog.SetDefault(slog.New(h))
}

func main() {
	flag.Parse()

	if *debugMode {
		slogLevel.Set(slog.LevelDebug)
	}

	s := server.NewMCPServer(
		"gist-mcp",
		"0.0.1",
		server.WithInstructions("A gist MCP server that allows you to create and manage gists"),
		server.WithToolCapabilities(false),
		server.WithResourceCapabilities(true, true),
		server.WithRecovery(),
	)

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		slog.Error("GITHUB_TOKEN environment variable is required")
		return
	}

	s.AddResources(resource.MustListGistResources(github.NewClient(nil).WithAuthToken(githubToken))...)
	s.AddPrompt(prompt.SummarizeGistPrompt())

	if err := server.ServeStdio(s); err != nil {
		slog.Error("Failed to start MCP server", "error", err)
		return
	}
}
