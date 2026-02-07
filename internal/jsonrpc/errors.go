package jsonrpc

// JSON-RPC 2.0 Error Codes
const (
	ErrorParse     = -32700 // Invalid JSON
	ErrorInvalidRequest = -32600 // Invalid Request
	ErrorMethodNotFound = -32601 // Method Not Found
	ErrorInvalidParams  = -32602 // Invalid Params
	ErrorInternal  = -32603 // Internal Error
)

// NewError creates a new JSON-RPC error
func NewError(code int, message string, data interface{}) *Error {
	return &Error{
		Code: code,
		Message: message,
		Data: data,
	}
}

// NewParseError creates a new JSON-RPC parse error
func NewParseError(message string, data interface{}) *Error {
	return NewError(ErrorParse, message, data)
}

// NewInvalidRequestError creates a new JSON-RPC invalid request error
func NewInvalidRequestError(message string, data interface{}) *Error {
	return NewError(ErrorInvalidRequest, message, data)
}

// NewMethodNotFoundError creates a new JSON-RPC method not found error
func NewMethodNotFoundError(message string, data interface{}) *Error {
	return NewError(ErrorMethodNotFound, message, data)
}

// NewInvalidParamsError creates a new JSON-RPC invalid params error
func NewInvalidParamsError(message string, data interface{}) *Error {
	return NewError(ErrorInvalidParams, message, data)
}

// NewInternalError creates a new JSON-RPC internal error
func NewInternalError(message string, data interface{}) *Error {
	return NewError(ErrorInternal, message, data)
}