package tool

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
)

// Calculator 结构体（仅用于组织方法）
type Calculator struct{}

// Name 定义工具名
func (c *Calculator) Name() string { return "calculator" }

// Description 定义工具描述
func (c *Calculator) Description() string {
	return "计算器工具，支持加减乘除运算"
}

// Execute 实现逻辑
func (c *Calculator) Execute(
	ctx context.Context,
	request mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {

	// 获取输入参数
	arguments := request.GetArguments()
	// 获取具体参数，并进行校验
	operation, ok := arguments["operation"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid operation argument")
	}
	a, ok := arguments["a"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid a argument")
	}
	b, ok := arguments["b"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid b argument")
	}

	// 要计算的结果
	var result float64
	// 执行计算
	switch operation {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	case "divide":
		if b == 0 {
			return nil, errors.New("不允许除以零")
		}
		result = a / b
	}

	log.Printf("运行结果：%v", result)

	// 返回结果
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("result: %v", result),
			},
		},
	}, nil
}
