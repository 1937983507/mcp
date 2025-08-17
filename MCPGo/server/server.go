package main

import (
	"bytes"
	"io"
	"log"
	"main/tool"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	// 服务类型
	serverType = "sse" // sse、streamable
	// 尾路径
	endpointPath = "mcp"
	// 端口
	port = ":8290"
)

func main() {

	// 开启日志记录
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 创建MCP服务器，添加工具
	mcpServer := newMCPServer()

	switch serverType {
	case "sse":
		// 创建并启动 SSE 服务器
		if err := startSSEServer(mcpServer); err != nil {
			log.Fatalf("服务启动失败：%v", err.Error())
			return
		}
	case "streamable":
		// 创建并启动 Streamable-HTTP 服务器
		if err := startStreamableServer(mcpServer); err != nil {
			log.Fatalf("服务启动失败：%v", err.Error())
			return
		}
	}

}

// 日志中间件
func loggingMiddleware(server http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[HTTP] %s %s %s from %s", r.Method, r.URL.Path, r.URL.RawQuery, r.RemoteAddr)

		// 如果是 POST，则读取 body
		if r.Method == http.MethodPost {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				log.Printf("[ERROR] 读取 Body 失败: %v", err)
			} else {
				log.Printf("[INFO] req Body: %s", string(bodyBytes))
				// 重新放回 Body（否则后续 handler 就读不到了）
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		server.ServeHTTP(w, r)
	})
}

// 创建MCP服务器，添加工具
func newMCPServer() *server.MCPServer {

	// 创建 MCP 服务器
	mcpServer := server.NewMCPServer(
		"MCP Server",
		"1.0.0",
		server.WithToolCapabilities(false),
		server.WithLogging(),
	)

	// 定义 calculator 工具
	calculator := &tool.Calculator{}
	calculatorSchema := mcp.NewTool(
		calculator.Name(),
		mcp.WithDescription(calculator.Description()),
		mcp.WithNumber("a", mcp.Required(), mcp.Description("参数1")),
		mcp.WithNumber("b", mcp.Required(), mcp.Description("参数2")),
		mcp.WithString("operation", mcp.Required(), mcp.Enum("add", "subtract", "multiply", "divide"), mcp.Description("运算符")),
	)

	// 将工具添加到 MCP 服务器
	mcpServer.AddTool(calculatorSchema, calculator.Execute)

	return mcpServer
}

// 创建并启动 SSE 服务器
func startSSEServer(mcpServer *server.MCPServer) error {

	// 创建 SSE 服务器
	sseServer := server.NewSSEServer(
		mcpServer,
		// server.WithBaseURL("http://127.0.0.1"),
		server.WithSSEEndpoint(endpointPath), // 设置 SSE endpoint path
		// server.WithKeepAlive(true),       // 保持长连接
		// server.WithKeepAliveInterval(10), // 保持长连接的时间
	)

	// // 启动服务
	// // Start 本质是用 net/http.ListenAndServe，所以可以直接用http包装，然后用中间件来获取HTTP层日志
	// if err := sseServer.Start(port); err != nil {
	// 	return err
	// }

	log.Println("启动服务")
	return http.ListenAndServe(port, loggingMiddleware(sseServer))
}

// 创建并启动 Streamable-HTTP 服务器
func startStreamableServer(mcpServer *server.MCPServer) error {

	// 创建 Streamable-HTTP 服务器
	streamableServer := server.NewStreamableHTTPServer(
		mcpServer,
		server.WithEndpointPath(endpointPath),
	)

	// // 启动服务
	// if err := streamableServer.Start(port); err != nil {
	// 	return err
	// }

	log.Println("启动服务")
	return http.ListenAndServe(port, loggingMiddleware(streamableServer))
}
