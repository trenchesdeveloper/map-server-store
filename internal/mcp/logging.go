package mcp

// ---- Logging ----

// LoggingLevel represents a severity level for log messages.
type LoggingLevel string

const (
	LogLevelDebug     LoggingLevel = "debug"
	LogLevelInfo      LoggingLevel = "info"
	LogLevelNotice    LoggingLevel = "notice"
	LogLevelWarning   LoggingLevel = "warning"
	LogLevelError     LoggingLevel = "error"
	LogLevelCritical  LoggingLevel = "critical"
	LogLevelAlert     LoggingLevel = "alert"
	LogLevelEmergency LoggingLevel = "emergency"
)

// SetLevelParams are sent by the client in a "logging/setLevel" request.
type SetLevelParams struct {
	Level LoggingLevel `json:"level"`
}

// LoggingMessageNotification is sent by the server as a "notifications/message" notification.
type LoggingMessageNotification struct {
	Level  LoggingLevel `json:"level"`
	Logger string       `json:"logger,omitempty"`
	Data   interface{}  `json:"data"`
}
