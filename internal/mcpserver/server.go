package mcpserver

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yusuke-hoguro/go-mcp-server/internal/tools/wikipedia"
)

const (
	serverName    = "blog-mcp-server"
	serverVersion = "0.1.0"
)

func New() *mcp.Server {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    serverName,
		Version: serverVersion,
	}, nil)

	wikipedia.Register(server)

	return server
}
