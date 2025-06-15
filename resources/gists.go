package resources

import (
	"context"
	"log/slog"
	"sort"
	"strings"

	"github.com/google/go-github/v72/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const RESOURCE_PREFIX = "gist://"

func RegisterGistResources(s *server.MCPServer, gh *github.Client) error {
	gists, err := listGists(context.Background(), gh) // maybe use a context with timeout
	if err != nil {
		return err
	}

	// keep only the most recent 50 gists
	var maxGists = 50
L:
	for _, gist := range gists {
		if maxGists <= 0 {
			break
		}

		for _, file := range gist.Files { // only add resources for files that are Markdown or tldraw
			if file.GetLanguage() != "Markdown" && !strings.HasSuffix(file.GetFilename(), ".tldraw") {
				continue L
			}
		}

		slog.Debug("gist info", "url", gist.GetHTMLURL(), "description", gist.GetDescription(), "updated_at", gist.GetUpdatedAt().String())
		s.AddResource(
			mcp.NewResource(
				RESOURCE_PREFIX+gist.GetID(),
				gist.GetDescription(), // TODO: maybe empty if no description
				mcp.WithMIMEType("application/json"),
			),
			handleReadGistResource(gh),
		)

		maxGists--
	}
	return nil
}

func listGists(ctx context.Context, gh *github.Client) ([]*github.Gist, error) {
	opt := &github.GistListOptions{
		ListOptions: github.ListOptions{
			PerPage: 100, // maximum allowed by GitHub API
		},
	}

	var allGists []*github.Gist
	for {
		gs, resp, err := gh.Gists.List(ctx, "", opt)
		if err != nil {
			return nil, err
		}

		allGists = append(allGists, gs...)
		slog.Debug("fetched gists", "count", len(gs), "next_page", resp.NextPage, "last_page", resp.LastPage)

		if resp.NextPage == 0 {
			break
		}

		opt.ListOptions.Page = resp.NextPage
	}

	sort.Slice(allGists, func(i, j int) bool {
		return allGists[i].UpdatedAt.Time.After(allGists[j].UpdatedAt.Time) // descending order by updated time
	})

	return allGists, nil
}

func handleReadGistResource(gh *github.Client) func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	return func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		gistID := strings.TrimPrefix(request.Params.URI, RESOURCE_PREFIX)
		gist, _, err := gh.Gists.Get(ctx, gistID)
		if err != nil {
			return nil, err
		}

		var sb strings.Builder
		for _, file := range gist.Files {
			sb.Write([]byte(file.GetContent()))
			sb.Write([]byte("\n\n")) // separate files with two newlines
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      request.Params.URI,
				MIMEType: "application/json",
				Text:     strings.TrimSuffix(sb.String(), "\n\n"),
			},
		}, nil
	}
}
