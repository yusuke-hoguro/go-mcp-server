package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yusuke-hoguro/go-mcp-server/internal/mcpserver"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))

	server := mcpserver.New()
	// 標準入出力を使ってMCPサーバーを起動する
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		logger.Error("mcp server stopped", slog.Any("error", err))
		os.Exit(1)
	}
}
