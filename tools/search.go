package tools

import (
	"context"
	"regexp"
	"strings"

	"github.com/CuteReimu/bilibili/v2"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func init() {
	Tools = append(Tools, server.ServerTool{
		Tool: mcp.NewTool("search_videos",
			mcp.WithDescription("Search for videos on Bilibili, China's largest online video site."),
			mcp.WithString("query",
				mcp.Required(),
				mcp.Description("search keyword"),
			),
		),
		Handler: toolSearchVideos,
	})
}

type SearchVideosResult struct {
	Videos []*VideoSearchInfo `json:"videos"`
}

type VideoSearchInfo struct {
	ID    int64    `json:"id"`
	Title string   `json:"title"`
	Url   string   `json:"url"`
	Tags  []string `json:"tags"`
	Cover string   `json:"cover"`
}

func toolSearchVideos(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query, err := request.RequireString("query")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := client.IntergratedSearch(bilibili.SearchParam{Keyword: query})
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	var videoResult *bilibili.SeachRespResult
	for _, result := range resp.Result {
		if result.ResultType == "video" {
			videoResult = &result
		}
	}
	if videoResult == nil {
		return mcp.NewToolResultJSON(&SearchVideosResult{})
	}

	videos := make([]*VideoSearchInfo, len(videoResult.Data))
	for i, data := range videoResult.Data {
		id, _ := convertor.ToInt(data["id"])
		title := convertor.ToString(data["title"])
		title = regexp.MustCompile(`<em class="keyword">(\S+)</em>`).ReplaceAllString(convertor.ToString(data["title"]), "$1")
		videos[i] = &VideoSearchInfo{
			ID:    id,
			Title: title,
			Url:   convertor.ToString(convertor.ToString(data["arcurl"])),
			Tags:  strings.Split(convertor.ToString(data["tag"]), ","),
			Cover: "https:" + convertor.ToString(data["pic"]),
		}
	}

	return mcp.NewToolResultJSON(&SearchVideosResult{Videos: videos})
}
