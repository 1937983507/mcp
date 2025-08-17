package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	httpURL       = "http://127.0.0.1:8290/mcp"
	transportType = "sse" // sse、streamable
)

func main() {

	// 开启日志记录
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create client based on transport type
	var c *client.Client
	var err error

	// 根据不同的传输类别，创建不同的客户端
	switch transportType {
	case "sse":
		// Create sse transport
		sseTransport, err := transport.NewSSE(httpURL)
		if err != nil {
			log.Fatalf("Failed to create sse transport: %v", err.Error())
		}
		// 创建客户端
		c = client.NewClient(sseTransport)
		// 启动客户端
		if err := c.Start(ctx); err != nil {
			log.Fatalf("Failed to start client: %v", err)
		}

	case "streamable":
		// Create HTTP transport
		httpTransport, err := transport.NewStreamableHTTP(httpURL)
		if err != nil {
			log.Fatalf("Failed to create HTTP transport: %v", err)
		}
		// Create client with the transport
		c = client.NewClient(httpTransport)
	}

	// Set up notification handler
	c.OnNotification(func(notification mcp.JSONRPCNotification) {
		log.Printf("Received notification: %s\n", notification.Method)
	})

	// 初始化客户端
	initRequest := mcp.InitializeRequest{
		Params: mcp.InitializeParams{
			ProtocolVersion: mcp.LATEST_PROTOCOL_VERSION,
			ClientInfo: mcp.Implementation{
				Name:    "MCP-Go Simple Client Example",
				Version: "1.0.0",
			},
			Capabilities: mcp.ClientCapabilities{},
		},
		Header: http.Header{},
	}
	serverInfo, err := c.Initialize(ctx, initRequest)
	if err != nil {
		log.Fatalf("Failed to initialize: %v", err)
	}

	// 展示服务器基本信息
	log.Printf("Connected to server: %s (version %s)\n",
		serverInfo.ServerInfo.Name,
		serverInfo.ServerInfo.Version)
	log.Printf("Server capabilities: %+v\n", serverInfo.Capabilities)

	// 列出服务器支持的可用工具
	if serverInfo.Capabilities.Tools != nil {
		toolsRequest := mcp.ListToolsRequest{}
		toolsResult, err := c.ListTools(ctx, toolsRequest)
		if err != nil {
			log.Printf("Failed to list tools: %v", err)
		} else {
			log.Printf("Server has %d tools available\n", len(toolsResult.Tools))
			for i, tool := range toolsResult.Tools {
				log.Printf("  %d. %s - %s\n", i+1, tool.Name, tool.Description)
			}
		}
	}

	// 列出服务器支持的可用资源
	if serverInfo.Capabilities.Resources != nil {
		resourcesRequest := mcp.ListResourcesRequest{}
		resourcesResult, err := c.ListResources(ctx, resourcesRequest)
		if err != nil {
			log.Printf("Failed to list resources: %v", err)
		} else {
			log.Printf("Server has %d resources available\n", len(resourcesResult.Resources))
			for i, resource := range resourcesResult.Resources {
				log.Printf("  %d. %s - %s\n", i+1, resource.URI, resource.Name)
			}
		}
	}

	// 开始调用某一工具
	callToolRequest := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name: "calculator",
			Arguments: map[string]any{
				"a":         1.2,
				"b":         2.2,
				"operation": "add",
			},
		},
	}
	callResult, err := c.CallTool(ctx, callToolRequest)
	if err != nil {
		log.Printf("Failed to call tools: %v", err)
	} else {
		log.Println("调用工具结果:", callResult.Content[0].(mcp.TextContent).Text)
	}

	log.Println("Client initialized successfully. Shutting down...")
	c.Close()
}
