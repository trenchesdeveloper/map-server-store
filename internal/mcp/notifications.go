package mcp

// ---- Notification Method Constants ----

const (
	// Server → Client notifications
	NotificationToolsListChanged     = "notifications/tools/listChanged"
	NotificationResourcesListChanged = "notifications/resources/listChanged"
	NotificationResourceUpdated      = "notifications/resources/updated"
	NotificationPromptsListChanged   = "notifications/prompts/listChanged"
	NotificationMessage              = "notifications/message"
	NotificationProgress             = "notifications/progress"

	// Client → Server notifications
	NotificationInitialized  = "notifications/initialized"
	NotificationCancelled    = "notifications/cancelled"
	NotificationRootsChanged = "notifications/roots/listChanged"
)

// ---- Method Constants ----

const (
	MethodInitialize = "initialize"
	MethodPing       = "ping"

	MethodToolsList = "tools/list"
	MethodToolsCall = "tools/call"

	MethodResourcesList         = "resources/list"
	MethodResourcesRead         = "resources/read"
	MethodResourcesTemplateList = "resources/templates/list"
	MethodResourcesSubscribe    = "resources/subscribe"
	MethodResourcesUnsubscribe  = "resources/unsubscribe"

	MethodPromptsList = "prompts/list"
	MethodPromptsGet  = "prompts/get"

	MethodLoggingSetLevel = "logging/setLevel"
)

// ---- Notification Payloads ----

// ResourceUpdatedNotification is sent by the server when a subscribed resource changes.
type ResourceUpdatedNotification struct {
	URI string `json:"uri"`
}

// ProgressNotification is sent by the server to report progress on a long-running request.
type ProgressNotification struct {
	ProgressToken interface{} `json:"progressToken"`
	Progress      float64     `json:"progress"`
	Total         float64     `json:"total,omitempty"`
}

// CancelledNotification is sent by the client to cancel an in-progress request.
type CancelledNotification struct {
	RequestID interface{} `json:"requestId"`
	Reason    string      `json:"reason,omitempty"`
}
