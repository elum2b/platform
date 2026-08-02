package mcp

import (
	"context"
	"fmt"

	"github.com/google/jsonschema-go/jsonschema"
	mcpserver "github.com/modelcontextprotocol/go-sdk/mcp"
)

type toolHandler[In, Out any] struct {
	tool        *mcpserver.Tool
	handler     mcpserver.ToolHandlerFor[In, Out]
	middleware  []func(context.Context) (context.Context, error)
	inputChecks []func(context.Context, any) error
}

func NewHandler[In, Out any](
	name string,
	description string,
	handler mcpserver.ToolHandlerFor[In, Out],
) Handler {
	successSchema, err := jsonschema.For[Out](nil)
	if err != nil {
		panic(fmt.Sprintf("mcp: create output schema for %q: %v", name, err))
	}

	errorSchema, err := jsonschema.For[ErrorResponse](nil)
	if err != nil {
		panic(fmt.Sprintf("mcp: create error schema for %q: %v", name, err))
	}

	return &toolHandler[In, Out]{
		tool: &mcpserver.Tool{
			Name:        name,
			Description: description,
			OutputSchema: &jsonschema.Schema{
				Type:  "object",
				AnyOf: []*jsonschema.Schema{successSchema, errorSchema},
			},
		},
		handler: handler,
	}
}

func (h *toolHandler[In, Out]) register(server *mcpserver.Server) {
	mcpserver.AddTool(server, h.tool, func(
		ctx context.Context,
		request *mcpserver.CallToolRequest,
		data In,
	) (*mcpserver.CallToolResult, any, error) {
		for _, middleware := range h.middleware {
			var err error

			ctx, err = middleware(ctx)
			if err != nil {
				return ErrorResult(h.tool.Name, err), nil, nil
			}
		}

		for _, check := range h.inputChecks {
			if err := check(ctx, data); err != nil {
				return ErrorResult(h.tool.Name, err), nil, nil
			}
		}

		result, output, err := h.handler(ctx, request, data)
		if err != nil {
			return ErrorResult(h.tool.Name, err), nil, nil
		}

		return result, output, nil
	})
}

func (h *toolHandler[In, Out]) Use(
	middleware func(context.Context) (context.Context, error),
) {
	h.middleware = append(h.middleware, middleware)
}

func (h *toolHandler[In, Out]) UseInput(
	check func(context.Context, any) error,
) {
	h.inputChecks = append(h.inputChecks, check)
}
