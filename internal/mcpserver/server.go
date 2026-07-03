package mcpserver

import "github.com/modelcontextprotocol/go-sdk/mcp"

const (
	serverName    = "blog-mcp-server"
	serverVersion = "0.1.0"
)

func New() *mcp.Server {
	return mcp.NewServer(&mcp.Implementation{
		Name:    serverName,
		Version: serverVersion,
	}, nil)
}
