package wikipedia

// 検索時の入力パラメータ
type SearchInput struct {
	Query    string `json:"query" jsonschema:"search query text"`
	Language string `json:"language,omitempty" jsonschema:"Wikipedia language code such as ja or en"`
	Limit    int    `json:"limit,omitempty" jsonschema:"maximum number of search results to return"`
}

// 全体の検索結果
type SearchOutput struct {
	Results []SearchResult `json:"results" jsonschema:"search results"`
}

// 1件の検索結果
type SearchResult struct {
	Title string `json:"title" jsonschema:"Wikipedia page title"`

	Snippet string `json:"snippet" jsonschema:"short search result snippet"`

	URL string `json:"url" jsonschema:"Wikipedia page URL"`
}
