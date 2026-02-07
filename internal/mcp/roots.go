package mcp

// ---- Roots ----

// Root represents a root directory that the client exposes to the server.
type Root struct {
	URI  string `json:"uri"`
	Name string `json:"name,omitempty"`
}

// ListRootsResult is returned by the client in response to "roots/list".
type ListRootsResult struct {
	Roots []Root `json:"roots"`
}
