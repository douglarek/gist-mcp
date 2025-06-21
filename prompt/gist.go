package prompt

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

// If you're using VS Code Copilot agent, you can invoke this prompt by typing a slash followed by the prompt name, like: /mcp.gist-mcp.gist_summary
func NewGistSummaryPrompt() mcp.Prompt {
	return mcp.NewPrompt(
		"gist_summary",
		mcp.WithPromptDescription("Summarize the contents of a Gist"),
	)
}

func HandleGistSummaryPrompt(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	return mcp.NewGetPromptResult("Gist summary",
		[]mcp.PromptMessage{
			mcp.NewPromptMessage(
				mcp.RoleUser,
				mcp.NewTextContent("What's the summary of this gist?"),
			),
		}), nil
}
