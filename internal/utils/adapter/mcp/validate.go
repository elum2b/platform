package mcp

import "github.com/go-playground/validator/v10"

var validate = validator.New()

// Validate checks MCP tool arguments against their validate tags.
func Validate(data any) bool {
	return validate.Struct(data) == nil
}
