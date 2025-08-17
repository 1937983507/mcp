# MCP

本仓库主要用于学习 **Model Context Protocol (MCP)** 的使用，包含两个不同框架实现的示例 **计算器 MCP 服务器**：

- **FastMCP**：基于 Python + [FastMCP](https://pypi.org/project/fastmcp/) 库实现
- **MCPGO**：基于 Go + [MCP-GO](https://github.com/mark3labs/mcp-go) 库实现

两者都实现了一个 **计算器工具 (calculator)**，支持 **加、减、乘、除** 四则运算，主要用于演示 MCP 协议在不同语言下的实现方式。

---

## 📂 仓库结构

```bash
mcp/
├── FastMCP/            # Python 版本
│ ├── server.py         # MCP 服务器（提供 calculator 工具）
│ ├── client_sse.py     # 客户端示例 1（使用 mcp.client.sse 连接）
│ └── client_sse2.py    # 客户端示例 2（使用 fastmcp.Client 连接）
│
└── MCPGO/          # Go 版本
├── server/         # 服务端
│ └── server.go
├── client/         # 客户端
│ └── client.go
└── tool/           # 工具实现
└── calculator.go
```

---

## 🚀 FastMCP (Python)

### 依赖安装

```bash
pip install fastmcp
```

### 运行服务器

```bash
cd FastMCP
python server.py
```

服务启动后将会输出如下，代表服务会在 http://127.0.0.1:8290/sse 提供 MCP 接口。

```bash


╭─ FastMCP 2.0 ──────────────────────────────────────────────────────────────╮
│                                                                            │
│        _ __ ___ ______           __  __  _____________    ____    ____     │
│       _ __ ___ / ____/___ ______/ /_/  |/  / ____/ __ \  |___ \  / __ \    │
│      _ __ ___ / /_  / __ `/ ___/ __/ /|_/ / /   / /_/ /  ___/ / / / / /    │
│     _ __ ___ / __/ / /_/ (__  ) /_/ /  / / /___/ ____/  /  __/_/ /_/ /     │
│    _ __ ___ /_/    \__,_/____/\__/_/  /_/\____/_/      /_____(_)____/      │
│                                                                            │
│                                                                            │
│                                                                            │
│    🖥️  Server name:     MCP Server Demo 🚀                                  │
│    📦 Transport:       SSE                                                 │
│    🔗 Server URL:      http://127.0.0.1:8290/sse                           │
│                                                                            │
│    📚 Docs:            https://gofastmcp.com                               │
│    🚀 Deploy:          https://fastmcp.cloud                               │
│                                                                            │
│    🏎️  FastMCP version: 2.11.3                                              │
│    🤝 MCP version:     1.13.0                                              │
│                                                                            │
╰────────────────────────────────────────────────────────────────────────────╯


[08/17/25 21:50:03] INFO     Starting MCP server 'MCP Server Demo 🚀' with transport 'sse' on http://127.0.0.1:8290/sse                                                                                                                                            server.py:1522
INFO:     Started server process [11848]
INFO:     Waiting for application startup.
INFO:     Application startup complete.
INFO:     Uvicorn running on http://127.0.0.1:8290 (Press CTRL+C to quit)
INFO:     127.0.0.1:14989 - "GET /sse HTTP/1.1" 200 OK
```

### 客户端示例 1

```bash
python client_sse.py
```

效果：初始化连接、获取工具列表并打印

```bash
[
    {
        "name": "calculator",
        "title": null,
        "description": null,
        "inputSchema": {
            "properties": {
                "a": {
                    "title": "A",
                    "type": "number"
                },
                "b": {
                    "title": "B",
                    "type": "number"
                },
                "operation": {
                    "title": "Operation",
                    "type": "string"
                }
            },
            "required": [
                "a",
                "b",
                "operation"
            ],
            "type": "object"
        },
        "outputSchema": {
            "properties": {
                "result": {
                    "title": "Result",
                    "type": "number"
                }
            },
            "required": [
                "result"
            ],
            "title": "_WrappedResult",
            "type": "object",
            "x-fastmcp-wrap-result": true
        },
        "annotations": null,
        "meta": {
            "_fastmcp": {
                "tags": []
            }
        }
    }
]
```

### 客户端示例 2

```bash
python client_sse2.py
```

效果：列出所有工具（包含 calculator）、调用 calculator 工具执行 1 + 2

示例输出：

```sql
Tool: calculator
Description: None
Parameters: {'properties': {'a': {'title': 'A', 'type': 'number'}, 'b': {'title': 'B', 'type': 'number'}, 'operation': {'title': 'Operation', 'type': 'string'}}, 'required': ['a', 'b', 'operation'], 'type': 'object'}
Tags: []
call tool result: CallToolResult(content=[TextContent(type='text', text='3.0', annotations=None, meta=None)], structured_content={'result': 3.0}, data=3.0, is_error=False)
```

### Cline

另外可以在支持 MCP 协议的客户端中使用（如 Cline/Cursor/RooCode 等）

需要在 mcp.json 配置中增加如下配置：

```bash
{
  "mcpServers": {
    "caculator":{
      "type":"sse",
      "url":"http://127.0.0.1:8290/sse"
    }
  }
}
```

---

## ⚡ MCPGO (Go)

### 依赖安装

```go
go get github.com/mark3labs/mcp-go/mcp
go get github.com/mark3labs/mcp-go/server
```

### 运行服务器

```bash
cd MCPGO
go run server/server.go
```

服务启动后将会输出如下，默认会启动在 http://127.0.0.1:8290/mcp .

```bash
PS xxxx\mcp\MCPGO> go run server/server.go
2025/08/17 21:57:45 server.go:114: 启动服务
2025/08/17 21:57:47 server.go:51: [HTTP] GET /sse  from 127.0.0.1:1237
```

### 运行客户端

```bash
cd MCPGO
go run client/client.go
```

效果：初始化连接、列出工具（calculator）、调用工具执行 1.2 + 2.2

示例输出：

```pgsql
PS D:\code\mcp\MCPGO> go run .\client\client.go
2025/08/17 21:59:03 client.go:80: Connected to server: MCP Server (version 1.0.0)
2025/08/17 21:59:03 client.go:83: Server capabilities: {Experimental:map[] Logging:0xf634c0 Prompts:<nil> Resources:<nil> Sampling:<nil> Tools:0xc000190bea}
2025/08/17 21:59:03 client.go:92: Server has 1 tools available
2025/08/17 21:59:03 client.go:94:   1. calculator - 计算器工具，支持加减乘除运算
2025/08/17 21:59:03 client.go:128: 调用工具结果: result: 3.4000000000000004
2025/08/17 21:59:03 client.go:131: Client initialized successfully. Shutting down...
```

### Cline

同样可以在支持 MCP 协议的客户端中使用（如 Cline/Cursor/RooCode 等）

需要在 mcp.json 配置中增加如下配置：

```bash
{
  "mcpServers": {
    "caculator":{
      "type":"sse",
      "url":"http://127.0.0.1:8290/mcp"
    }
  }
}
```

---

## 📖 学习要点

1. **MCP 协议** 提供了标准化的 Server/Client 交互方式

2. **FastMCP** 适合快速验证和原型开发

3. **MCPGO** 更接近工业级应用，结构化清晰，支持扩展更多工具

4. 两个实现展示了 **跨语言一致性** —— 不论 Python 还是 Go，都能实现相同的功能，并通过 MCP 协议通信

---

## 🔮 后续扩展方向

添加更多工具（如字符串处理、文件操作等）

将 MCP 工具与 Agent 进行对接

---

## ❓ FAQ / 常见问题

**Q1:** 客户端连接失败怎么办？

- 检查服务器是否启动，并确认端口号与 URL 匹配
- SSE 协议默认端口 8290，可自定义

**Q2:** Python 安装依赖失败？

- 建议使用 `python -m pip install --upgrade pip` 更新 pip
- 确认使用 Python 3.9+

**Q3:** Go 编译报错？

- 确认 go.mod 已初始化
- 运行 `go mod tidy` 下载依赖

**Q4:** 除以零会如何处理？

- FastMCP：抛出 `ValueError`
- MCPGO：返回 error 信息

---

## 📌 参考

[fastmcp](https://github.com/jlowin/fastmcp)

[mcp-go](https://github.com/mark3labs/mcp-go)

[Model Context Protocol (MCP)](https://modelcontextprotocol.io/overview)
