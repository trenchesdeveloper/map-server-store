package mcp

// ---- Resources ----

// Resource represents a known resource that the server can read.
type Resource struct {
	URI         string `json:"uri"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
}

// ResourceTemplate describes a URI template for dynamic resources.
type ResourceTemplate struct {
	URITemplate string `json:"uriTemplate"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
}

// ---- Resource List ----

// ListResourcesParams are sent by the client in a "resources/list" request.
type ListResourcesParams struct {
	PaginatedRequest
}

// ListResourcesResult is returned by "resources/list".
type ListResourcesResult struct {
	Resources []Resource `json:"resources"`
	PaginatedResult
}

// ListResourceTemplatesParams are sent by the client in a "resources/templates/list" request.
type ListResourceTemplatesParams struct {
	PaginatedRequest
}

// ListResourceTemplatesResult is returned by "resources/templates/list".
type ListResourceTemplatesResult struct {
	ResourceTemplates []ResourceTemplate `json:"resourceTemplates"`
	PaginatedResult
}

// ---- Resource Read ----

// ReadResourceParams are sent by the client in a "resources/read" request.
type ReadResourceParams struct {
	URI string `json:"uri"`
}

// ReadResourceResult is returned by "resources/read".
type ReadResourceResult struct {
	Contents []ResourceContents `json:"contents"`
}

// ResourceContents holds the content of a single resource.
// Either Text or Blob (base64) will be populated, but not both.
type ResourceContents struct {
	URI      string `json:"uri"`
	MimeType string `json:"mimeType,omitempty"`
	Text     string `json:"text,omitempty"` // for text resources
	Blob     string `json:"blob,omitempty"` // for binary resources (base64)
}

// ---- Resource Subscriptions ----

// SubscribeParams are sent by the client in a "resources/subscribe" request.
type SubscribeParams struct {
	URI string `json:"uri"`
}

// UnsubscribeParams are sent by the client in a "resources/unsubscribe" request.
type UnsubscribeParams struct {
	URI string `json:"uri"`
}
