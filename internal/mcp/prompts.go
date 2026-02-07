package mcp

// ---- Prompts ----

// Prompt describes an MCP prompt the server exposes.
type Prompt struct {
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Arguments   []PromptArgument `json:"arguments,omitempty"`
}

// PromptArgument describes a single argument that a prompt accepts.
type PromptArgument struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

// ---- Prompt List ----

// ListPromptsParams are sent by the client in a "prompts/list" request.
type ListPromptsParams struct {
	PaginatedRequest
}

// ListPromptsResult is returned by "prompts/list".
type ListPromptsResult struct {
	Prompts []Prompt `json:"prompts"`
	PaginatedResult
}

// ---- Prompt Get ----

// GetPromptParams are sent by the client in a "prompts/get" request.
type GetPromptParams struct {
	Name      string            `json:"name"`
	Arguments map[string]string `json:"arguments,omitempty"`
}

// GetPromptResult is returned by "prompts/get".
type GetPromptResult struct {
	Description string          `json:"description,omitempty"`
	Messages    []PromptMessage `json:"messages"`
}

// PromptMessage represents a single message within a prompt result.
type PromptMessage struct {
	Role    string  `json:"role"` // "user" or "assistant"
	Content Content `json:"content"`
}
