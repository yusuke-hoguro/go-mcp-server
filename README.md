# go-mcp-server
Go で MCP サーバーを作成するためのリポジトリです。

## 目的

最終的には自作 Blog API と連携できる Go 製 MCP サーバーを作成します。
最初の POC として、Wikipedia 検索 Tool を実装します。

## 現在のステップ

Step 1 では、stdio で起動できる最小構成の MCP サーバーを作成します。

```sh
go run ./cmd/blog-mcp-server
```

現在のサーバーにはまだ Tool を登録していません。
次のステップでは、`wikipedia_search` Tool のスキーマと handler の境界を定義します。
