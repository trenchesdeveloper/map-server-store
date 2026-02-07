package mcp

import "encoding/json"

// MCP Protocol version
const ProtocolVersion = "2024-11-05"

// ---- Capability types ----

// ServerCapabilities describes what the MCP server supports.
type ServerCapabilities struct {
	Tools     *ToolCapability     `json:"tools,omitempty"`
	Resources *ResourceCapability `json:"resources,omitempty"`
	Prompts   *PromptCapability   `json:"prompts,omitempty"`
	Logging   *LoggingCapability  `json:"logging,omitempty"`
}

type ToolCapability struct {
	ListChanged bool `json:"listChanged,omitempty"`
}

type ResourceCapability struct {
	Subscribe   bool `json:"subscribe,omitempty"`
	ListChanged bool `json:"listChanged,omitempty"`
}

type PromptCapability struct {
	ListChanged bool `json:"listChanged,omitempty"`
}

type LoggingCapability struct{}

// ClientCapabilities describes what the MCP client supports.
type ClientCapabilities struct {
	Roots    *RootsCapability    `json:"roots,omitempty"`
	Sampling *SamplingCapability `json:"sampling,omitempty"`
}

type RootsCapability struct {
	ListChanged bool `json:"listChanged,omitempty"`
}

type SamplingCapability struct{}

// ---- Initialize ----

// InitializeParams are sent by the client in the "initialize" request.
type InitializeParams struct {
	ProtocolVersion string             `json:"protocolVersion"`
	Capabilities    ClientCapabilities `json:"capabilities"`
	ClientInfo      Implementation     `json:"clientInfo"`
}

// InitializeResult is returned by the server in response to "initialize".
type InitializeResult struct {
	ProtocolVersion string             `json:"protocolVersion"`
	Capabilities    ServerCapabilities `json:"capabilities"`
	ServerInfo      Implementation     `json:"serverInfo"`
	Instructions    string             `json:"instructions,omitempty"`
}

// Implementation identifies a client or server.
type Implementation struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ---- Tools ----

// Tool describes an MCP tool the server exposes.
type Tool struct {
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	InputSchema json.RawMessage `json:"inputSchema"`
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

// ToolListResult is returned by "tools/list".
type ToolListResult struct {
	Tools []Tool `json:"tools"`
}

// ---- Content types ----

// Content represents a content block in an MCP response.
type Content struct {
	Type     string `json:"type"`               // "text", "image", "resource"
	Text     string `json:"text,omitempty"`     // for type "text"
	MimeType string `json:"mimeType,omitempty"` // for type "image"
	Data     string `json:"data,omitempty"`     // for type "image" (base64)
	URI      string `json:"uri,omitempty"`      // for type "resource"
}

// NewTextContent creates a text content block.
func NewTextContent(text string) Content {
	return Content{
		Type: "text",
		Text: text,
	}
}

// NewErrorContent creates a text content block marked as an error.
func NewErrorContent(text string) (Content, bool) {
	return Content{
		Type: "text",
		Text: text,
	}, true
}

// ---- Ping ----

// PingResult is returned by the server in response to "ping".
type PingResult struct{}
