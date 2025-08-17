import asyncio
from fastmcp import Client

# # 配置 MCP 服务器地址
# config = {
#     "mcpServers": {
#         "printer": {
#             "url": "http://127.0.0.1:8000/sse",
#             "transport": "sse"
#         }
#     }
# }
#
# # 初始化客户端
# client = Client(config)

# 初始化客户端
client = Client("http://127.0.0.1:8290/sse")

async def main():
    async with client:

        # 获取所有可用工具
        tools = await client.list_tools()
        for tool in tools:
            print(f"Tool: {tool.name}")
            print(f"Description: {tool.description}")
            if tool.inputSchema:
                print(f"Parameters: {tool.inputSchema}")
            # Access tags and other metadata
            if hasattr(tool, 'meta') and tool.meta:
                fastmcp_meta = tool.meta.get('_fastmcp', {})
                print(f"Tags: {fastmcp_meta.get('tags', [])}")

        # 调用工具
        result = await client.call_tool("calculator", {"a": 1, "b": 2, "operation":"add"})
        print(f"call tool result: {result}")

if __name__ == "__main__":
    asyncio.run(main())