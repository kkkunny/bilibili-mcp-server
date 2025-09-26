package main

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/kkkunny/bilibili-mcp-server/tools"
)

func main() {
	mapSvr := server.NewMCPServer("BiliBili", "0.0.1", server.WithLogging(), server.WithRecovery())
	mapSvr.AddTools(tools.Tools...)
	httpSvr := server.NewStreamableHTTPServer(mapSvr)
	if err := httpSvr.Start(":80"); err != nil {
		panic(err)
	}
}
