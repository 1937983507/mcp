# server.py
from fastmcp import FastMCP

mcp = FastMCP("MCP Server Demo 🚀")

@mcp.tool()
def calculator(a: float, b: float, operation: str) -> float:
    if operation=="add":
        """Add two numbers"""
        return a+b
    elif operation=="subtract":
        """Return a - b"""
        return a-b
    elif operation=="multiply":
        """Return a * b"""
        return a*b
    elif operation=="divide":
        """Return a / b"""
        if b == 0:
            raise ValueError("Division by zero")
        return a / b       

if __name__ == "__main__":
    mcp.run(transport="sse", host="127.0.0.1", port=8290)
