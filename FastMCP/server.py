# server.py
from fastmcp import FastMCP

mcp = FastMCP("Demo 🚀")

@mcp.tool()
def add(a: float, b: float) -> float:
    """Add two numbers"""
    return a + b

@mcp.tool()
def sub(a: float, b: float) -> float:
    """Return a - b"""
    return a - b

@mcp.tool()
def mul(a: float, b: float) -> float:
    """Return a * b"""
    return a * b

@mcp.tool()
def div(a: float, b: float) -> float:
    """Return a / b"""
    if b == 0:
        raise ValueError("Division by zero")
    return a / b

if __name__ == "__main__":
    mcp.run(transport="sse", host="127.0.0.1", port=8000)
