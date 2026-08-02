package adapter

import (
	"context"

	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"
	"github.com/gofiber/fiber/v3"
	mcpserver "github.com/modelcontextprotocol/go-sdk/mcp"

	httputils "github.com/elum2b/platform/internal/utils/adapter/http"
	mcputils "github.com/elum2b/platform/internal/utils/adapter/mcp"
	socketutils "github.com/elum2b/platform/internal/utils/adapter/socket"
)

// Registry contains the API transports available while registering methods.
type Registry struct {
	HTTP   fiber.Router
	Socket etp.Router
	MCP    mcputils.Router
}

// Register exposes a method through every enabled transport in the registry.
func (method Method[In, Out]) Register(registry Registry) {
	if method.Transports&HTTP != 0 && registry.HTTP != nil {
		method.registerHTTP(registry.HTTP)
	}

	if method.Transports&WS != 0 && registry.Socket != nil {
		method.registerSocket(registry.Socket)
	}

	if method.Transports&MCP != 0 && registry.MCP != nil {
		method.registerMCP(registry.MCP)
	}
}

func (method Method[In, Out]) registerHTTP(router fiber.Router) {
	handler := func(ctx fiber.Ctx) error {
		data := new(In)
		if !httputils.Decode(ctx, data) {
			return httputils.Error(ctx, serviceerrors.ErrInvalidFields)
		}

		response, err := method.call(&Context{
			Context:   ctx.Context(),
			Transport: HTTP,
			HTTP:      ctx,
		}, *data)
		if err != nil {
			return httputils.Error(ctx, err)
		}

		return httputils.Respond(ctx, response)
	}

	switch method.Method {
	case fiber.MethodGet:
		router.Get(method.Key, handler)
	case fiber.MethodPost:
		router.Post(method.Key, handler)
	case "":
		router.Get(method.Key, handler)
		router.Post(method.Key, handler)
	default:
		panic("adapter: unsupported HTTP method")
	}
}

func (method Method[In, Out]) registerSocket(router etp.Router) {
	router.On(method.Key, func(ctx *etp.Context) error {
		data := new(In)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		response, err := method.call(&Context{
			Context:   ctx,
			AccountID: socketAccountID(ctx),
			Transport: WS,
			Socket:    ctx,
		}, *data)
		if err != nil {
			return err
		}

		return socketutils.Respond(ctx, method.Key, response)
	})
}

func (method Method[In, Out]) registerMCP(router mcputils.Router) {
	router.On(method.Key, mcputils.NewHandler(
		method.Key,
		method.Description,
		func(
			ctx context.Context,
			_ *mcpserver.CallToolRequest,
			data In,
		) (*mcpserver.CallToolResult, Out, error) {
			if !mcputils.Validate(data) {
				var empty Out

				return nil, empty, serviceerrors.ErrInvalidFields
			}

			principal, _ := mcputils.PrincipalFromContext(ctx)

			response, err := method.call(&Context{
				Context:   ctx,
				AccountID: principal.AccountID,
				Transport: MCP,
			}, data)
			if err != nil {
				var empty Out

				return nil, empty, err
			}

			return nil, response, nil
		},
	))
}

func socketAccountID(ctx *etp.Context) string {
	if ctx == nil || ctx.Peer == nil {
		return ""
	}

	return ctx.Peer.Identity().UserID
}
