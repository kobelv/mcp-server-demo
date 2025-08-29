package adapter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// NewMCPServer 创建一个新的 MCPServer 实例。
//
// 这个函数首先创建一个新的 MCPServer 对象，使用指定的名称（"demo"）、版本（"0.1.0"）和配置（此处禁用了工具能力）。
//
// 接着，它定义了一个新的工具（"greeting"），该工具接受一个名为 "name" 的字符串参数，该参数是必需的，并附带了描述信息（"name of the person"）。
// 工具还附带了一个描述（"greet to the input name"）。
//
// 然后，将定义的工具添加到 MCPServer 实例中，并关联了一个处理函数 GreetHandler，该函数用于处理该工具的请求。
//
// 最后，返回创建好的 MCPServer 实例。
func NewMCPServer() *server.MCPServer {
	s := server.NewMCPServer("mcp-server-demo", "0.1.0", server.WithToolCapabilities(false))

	t := mcp.NewTool("hellow_word",
		mcp.WithDescription("this tool is to greet to the input name"),
		mcp.WithString("name", mcp.Required(), mcp.Description("name from the input")),
	)
	s.AddTool(t, GreetHandler)

	t2 := mcp.NewTool("get_current_date",
		mcp.WithDescription("this tool is to get current date"))
	s.AddTool(t2, GetCurrentDateHandler)

	return s
}

// GreetHandler 是一个处理请求的函数，它接收一个上下文（context.Context）和一个 CallToolRequest 类型的请求参数，
// 返回一个 CallToolResult 类型的指针和一个错误（error）。
//
// 参数：
// - ctx: 上下文对象，用于传递请求范围内的值、取消信号、截止日期等。
// - req: CallToolRequest 类型的请求参数，包含了工具调用的请求信息。
//
// 返回值：
// - *mcp.CallToolResult: 指向 CallToolResult 的指针，包含了工具调用的结果。如果发生错误，该值可能为 nil。
// - error: 如果函数执行过程中发生错误，将返回一个非 nil 的错误对象；否则返回 nil。
//
// 函数逻辑：
// 1. 从 req.GetArguments() 中获取名为 "name" 的参数值，并尝试将其断言为字符串类型。
// 2. 如果 "name" 参数不存在或不是字符串类型，函数将返回一个错误，指出 "name must be a string"。
// 3. 如果 "name" 参数有效，函数将使用 fmt.Sprintf 格式化一个包含问候语的字符串，并通过 mcp.NewToolResultText 创建一个新的 CallToolResult 对象。
// 4. 返回创建的 CallToolResult 对象和一个 nil 错误值。
func GreetHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, ok := req.GetArguments()["name"].(string)
	if !ok {
		return nil, errors.New("name must be a string")
	}
	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}

func GetCurrentDateHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText(time.Now().Format("2006/01/02")), nil
}
