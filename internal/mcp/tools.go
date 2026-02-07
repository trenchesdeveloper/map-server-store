package mcp

// ---- Tools ----

// Tool describes an MCP tool the server exposes.
type Tool struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	InputSchema InputSchema `json:"inputSchema"`
}

type InputSchema struct {
	Type       string            `json:"type"`
	Properties map[string]string `json:"properties"`
	Required   []string          `json:"required,omitempty"`
}

// ToolListParams are sent by the client in a "tools/list" request.
type ToolListParams struct {
	PaginatedRequest
}

// ToolListResult is returned by "tools/list".
type ToolListResult struct {
	Tools []Tool `json:"tools"`
	PaginatedResult
}

// ToolCallParams are sent by the client in a "tools/call" request.
type ToolCallParams struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments,omitempty"`
}

// ToolCallResult is returned by the server after executing a tool.
type ToolCallResult struct {
	Content []Content `json:"content"`
	IsError bool      `json:"isError,omitempty"`
}
