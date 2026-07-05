package wikipedia

import (
	"context"
	"testing"
)

func TestSearchReturnsPlaceholderResult(t *testing.T) {
	_, output, err := Search(context.Background(), nil, SearchInput{
		Query:    "Model Context Protocol",
		Language: "en",
		Limit:    3,
	})
	if err != nil {
		t.Fatalf("Search returned error: %v", err)
	}

	if len(output.Results) != 1 {
		t.Fatalf("len(output.Results) = %d, want 1", len(output.Results))
	}

	result := output.Results[0]
	if result.Title == "" {
		t.Fatal("result.Title is empty")
	}
	if result.Snippet == "" {
		t.Fatal("result.Snippet is empty")
	}

	wantURL := "https://en.wikipedia.org/wiki/Special:Search?search=Model+Context+Protocol"
	if result.URL != wantURL {
		t.Fatalf("result.URL = %q, want %q", result.URL, wantURL)
	}
}
