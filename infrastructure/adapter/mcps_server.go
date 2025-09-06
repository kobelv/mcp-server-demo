package adapter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// NewMCPServer 创建一个新的 MCPServer 实例，添加了2个工具。
func NewMCPServer() *server.MCPServer {
	s := server.NewMCPServer("mcp-server-demo", "0.1.0", server.WithToolCapabilities(false), server.WithRecovery())

	t := mcp.NewTool("hello_world",
		mcp.WithDescription("say hello to someone"),
		mcp.WithString("greet_name", mcp.Required(), mcp.Description("name from the person to greet")),
		mcp.WithString("greet_message", mcp.Description("message to greet"), mcp.DefaultString("have a good day")),
	)
	s.AddTool(t, GreetHandler)

	return s
}

// GreetHandler 是一个请求hello_world工具时的处理函数
func GreetHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	greet_name, err := req.RequireString("greet_name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), errors.New("greet_name must be a string")
	}

	greet_message, _ := req.GetArguments()["greet_message"].(string)

	return mcp.NewToolResultText(fmt.Sprintf("你好 %s, 今天是 %s, %s", greet_name, time.Now().Format("2006/01/02"), greet_message)), nil
}
