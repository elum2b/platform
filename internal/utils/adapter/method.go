package adapter

import (
	"context"

	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/internalapi"
	serviceerrors "github.com/elum2b/services/errors"
	"github.com/go-playground/validator/v10"
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/services"
)

// Transport identifies an API transport that can expose a method.
type Transport uint8

const (
	HTTP Transport = 1 << iota
	WS
	MCP
)

// Context contains transport-neutral data available to a method.
type Context struct {
	Context   context.Context
	AccountID string
	Transport Transport
	Data      any
	HTTP      fiber.Ctx
	Socket    *etp.Context
}

// Middleware runs before a method handler.
type Middleware func(*Context) error

// Method defines one API method that can be exposed through multiple transports.
type Method[In, Out any] struct {
	Key         string
	Description string
	Transports  Transport
	Method      string
	Middleware  []Middleware
	Handler     func(*Context, In) (Out, error)
}

func (method Method[In, Out]) call(ctx *Context, data In) (Out, error) {
	ctx.Data = data

	for _, middleware := range method.Middleware {
		if err := middleware(ctx); err != nil {
			var empty Out

			return empty, err
		}
	}

	return method.Handler(ctx, data)
}

// WorkspaceAccess requires the account to access a workspace method.
func WorkspaceAccess(method string) Middleware {
	return func(ctx *Context) error {
		if ctx.AccountID == "" {
			return serviceerrors.ErrUnauthorized
		}

		data := new(workspaceAccessRequest)
		if !Decode(ctx.Data, data) {
			return serviceerrors.ErrInvalidFields
		}

		allowed, err := services.Control.Internal.CheckWorkspaceAccess(
			ctx.Context,
			internalapi.WorkspaceAccessRequest{
				AccountID:   ctx.AccountID,
				WorkspaceID: data.WorkspaceID,
				MethodKey:   method,
			},
		)
		if err != nil {
			return err
		}

		if !allowed {
			return serviceerrors.ErrForbidden
		}

		return nil
	}
}

// GlobalAccess requires the account to access a global method.
func GlobalAccess(method string) Middleware {
	return func(ctx *Context) error {
		if ctx.AccountID == "" {
			return serviceerrors.ErrUnauthorized
		}

		allowed, err := services.Control.Internal.CheckGlobalAccess(
			ctx.Context,
			internalapi.GlobalAccessRequest{
				AccountID: ctx.AccountID,
				MethodKey: method,
			},
		)
		if err != nil {
			return err
		}

		if !allowed {
			return serviceerrors.ErrForbidden
		}

		return nil
	}
}

type workspaceAccessRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

var validate = validator.New()

// Decode converts method input into a middleware request and validates it.
func Decode(data any, target any) bool {
	encoded, err := json.Marshal(data)
	if err != nil {
		return false
	}

	if err := json.Unmarshal(encoded, target); err != nil {
		return false
	}

	return validate.Struct(target) == nil
}
