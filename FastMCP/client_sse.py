import asyncio
import json

from mcp import ClientSession
from mcp.client.sse import sse_client

async def run():
    sse_url = "http://127.0.0.1:8000/sse"
    async with sse_client(url=sse_url) as streams, ClientSession(*streams) as session:
        # 初始化
        await session.initialize()

        # 获取 tool 列表
        response = await session.list_tools()
        print(json.dumps([tool.model_dump() for tool in response.tools], indent=4, ensure_ascii=False))

if __name__ == "__main__":
    asyncio.run(run())







