# 📓 Gist MCP

A golang rewrite of [GistPad MCP](https://github.com/lostintangent/gistpad-mcp) using the [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go) library.

## Usage

Before building, you must install go 1.24+ first.

```bash
git clone https://github.com/douglarek/gist-mcp.git
cd gist-mcp
go build -o gist-mcp ./cmd/gist-mcp/main.go
```

### Cursor Editor Integration

To use this server in Cursor, you can add the following to your `mcp.json` file:

```json
{
  "mcpServers": {
    "unsplash": {
      "command": "<gist-mcp binary dir>",
      "args": [],
      "env": {
        "GITHUB_TOKEN": "<your-github-token>"
      }
    }
  }
}
```

### VS Code Editor Integration

```json
  "mcp": {
    "inputs": [
      {
        "type": "promptString",
        "id": "github-token",
        "description": "GitHub Token",
        "password": true
      }
    ],
    "servers": {
      "gist-mcp": {
        "type": "stdio",
        "command": "/home/douglarek/bin/gist-mcp",
        // "args": [
        //   "--debug"
        // ],
        "env": {
          "GITHUB_TOKEN": "${input:github-token}"
        }
      }
    }
  }
```

VS Code will prompt you for the github token when the server starts.

## Why rewrite

Due to resource occupation and dependency issues, the Go implementation only needs to provide an executable binary.