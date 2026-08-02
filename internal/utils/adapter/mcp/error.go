package mcp

import (
	"fmt"

	serviceerrors "github.com/elum2b/services/errors"
	mcpserver "github.com/modelcontextprotocol/go-sdk/mcp"
)

// ErrorResponse contains the structured error returned by an MCP tool.
type ErrorResponse struct {
	Error ErrorData `json:"error"`
}

// ErrorData contains the stable error key and its public message.
type ErrorData struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

// ErrorResult converts a service error into an MCP tool error result.
func ErrorResult(method string, err error) *mcpserver.CallToolResult {
	key := serviceerrors.CodeOf(err)
	message := serviceerrors.MessageOf(err)

	if key == "" {
		key = serviceerrors.CodeInternalError
	}

	if message == "" {
		message = "internal error"
	}

	return &mcpserver.CallToolResult{
		Content: []mcpserver.Content{
			&mcpserver.TextContent{Text: errorInstruction(method, key)},
		},
		StructuredContent: ErrorResponse{
			Error: ErrorData{Key: key, Message: message},
		},
		IsError: true,
	}
}

func errorInstruction(method, key string) string {
	switch key {
	case serviceerrors.CodeUnauthorized:
		return "The MCP token has expired. Ask the user to issue a new MCP " +
			"token before retrying."
	case serviceerrors.CodeForbidden:
		return fmt.Sprintf(
			"The account does not have permission to call %s.",
			method,
		)
	case serviceerrors.CodeInvalidFields:
		return fmt.Sprintf(
			"The arguments for %s are invalid. Correct them before retrying.",
			method,
		)
	case serviceerrors.CodeNotFound:
		return "The requested resource was not found."
	case serviceerrors.CodeConflict:
		return "The requested operation conflicts with the current state."
	case serviceerrors.CodeFailedPrecondition:
		return "The operation cannot be completed in the current state."
	case serviceerrors.CodeNotReady, serviceerrors.CodeUnavailable:
		return "The service is temporarily unavailable. Retry later."
	case serviceerrors.CodeTimeout:
		return "The operation timed out. Retry later."
	case serviceerrors.CodeRateLimit:
		return "The rate limit was exceeded. Retry later."
	default:
		return "The method could not be completed because of an internal error."
	}
}
