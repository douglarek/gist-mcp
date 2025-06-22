package prompt

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// SummarizeGistPrompt returns a new MCP prompt for summarizing a Gist.
// If you're using VS Code Copilot agent, you can invoke this prompt by typing a slash followed by the prompt name, like: /mcp.gist-mcp.gist_summary
func SummarizeGistPrompt() (mcp.Prompt, server.PromptHandlerFunc) {
	return mcp.NewPrompt(
			"gist_summary",
			mcp.WithPromptDescription("Summarize the contents of a Gist"),
		), func(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			return mcp.NewGetPromptResult("Gist summary",
				[]mcp.PromptMessage{
					mcp.NewPromptMessage(
						mcp.RoleUser,
						mcp.NewTextContent("What's the summary of this gist?"),
					),
				}), nil
		}
}
