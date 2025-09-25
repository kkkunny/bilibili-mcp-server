package tools

import (
	"os"

	"github.com/CuteReimu/bilibili/v2"
	"github.com/mark3labs/mcp-go/server"
)

var client = bilibili.New()

func init() {
	cookie := os.Getenv("BILIBILI_COOKIE")
	client.SetCookiesString(cookie)
}

var Tools []server.ServerTool
