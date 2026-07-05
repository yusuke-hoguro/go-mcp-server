package wikipedia

import (
	"context"
	"fmt"
	"net/url"
	"regexp"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const ToolName = "wikipedia_search"

var languageCodePattern = regexp.MustCompile(`^[a-z]{2,3}$`)

// MCPクライアントが操作できるツールを登録する
func Register(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        ToolName,
		Description: "Search Wikipedia articles by keyword.",
	}, Search)
}

// Wikipedia検索ツールの実装 (Placeholder)
// context.Contextとmcp.CallToolRequestはMCPの仕様上必要だが、ここでは使用しない
func Search(_ context.Context, _ *mcp.CallToolRequest, input SearchInput) (
	*mcp.CallToolResult,
	SearchOutput,
	error,
) {
	language := input.Language
	if language == "" {
		language = "ja"
	}
	// Wikipediaの言語コードが不正な場合はエラーを返す
	if !languageCodePattern.MatchString(language) {
		return nil, SearchOutput{}, fmt.Errorf("invalid Wikipedia language code: %q", input.Language)
	}

	return nil, SearchOutput{
		Results: []SearchResult{
			{
				Title:   "Wikipedia search placeholder",
				Snippet: "The wikipedia_search tool boundary is registered. Wikipedia API access will be implemented in a later step.",
				URL:     "https://" + language + ".wikipedia.org/wiki/Special:Search?search=" + url.QueryEscape(input.Query),
			},
		},
	}, nil
}
