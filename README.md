# go-mcp-server
Go で MCP サーバーを作成するためのリポジトリです。

## 目的

最終的には自作 Blog API と連携できる Go 製 MCP サーバーを作成します。
最初の POC として、Wikipedia 検索 Tool を実装します。

## 現在のステップ

Step 2 では、`wikipedia_search` Tool のスキーマと handler の境界を定義します。

```sh
go run ./cmd/blog-mcp-server
```

現在の `wikipedia_search` Tool は境界確認用のダミー結果を返します。
実際の Wikipedia API 呼び出しは、次のステップで `internal/clients/wikipedia`
として実装します。
