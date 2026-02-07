package jsonrpc

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

// Handler is a function that handles a JSON-RPC request and returns a result or error.
type Handler func(params json.RawMessage) (interface{}, *Error)

// Server is a JSON-RPC 2.0 server that reads requests from an io.Reader
// and writes responses to an io.Writer (typically stdin/stdout for stdio transport).
type Server struct {
	handlers map[string]Handler
	logger   *logrus.Logger
}

// NewServer creates a new JSON-RPC server with the given reader and writer.
func NewServer(logger *logrus.Logger) *Server {
	return &Server{
		handlers: make(map[string]Handler),
		logger:   logger,
	}
}

func (s *Server) RegisterMethod(method string, handler Handler) {
	s.handlers[method] = handler
	s.logger.WithField("method", method).Info("Registered method")
}

func (s *Server) HandleRequest(req *Request) *Response {
	s.logger.WithFields(logrus.Fields{
		"method": req.Method,
		"id":     req.ID,
	}).Debug("Handling request")

	if err := req.Validate(); err != nil {
		var jsonErr *Error
		if errors.As(err, &jsonErr) {
			return NewErrorResponse(req.ID, jsonErr)
		}
		return NewErrorResponse(req.ID, NewInternalError("Internal error", err))
	}

	handler, ok := s.handlers[req.Method]
	if !ok {
		return NewErrorResponse(req.ID, NewMethodNotFoundError(
			fmt.Sprintf("Method '%s' not found", req.Method), nil,
		))
	}

	result, err := handler(req.Params)
	if err != nil {
		var jsonErr *Error
		if errors.As(err, &jsonErr) {
			return NewErrorResponse(req.ID, jsonErr)
		}
		return NewErrorResponse(req.ID, NewInternalError("Internal error", err))
	}
	return NewSuccessResponse(req.ID, result)
}

