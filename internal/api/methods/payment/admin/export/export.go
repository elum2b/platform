package exportapi

import (
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type Request struct {
	WorkspaceID string    `json:"workspace_id"  validate:"required,uuid"`
	Now         time.Time `json:"now,omitempty"`
}

type Response struct {
	Package padm.ExportPackage `json:"package"`
}

var (
	methodKey         = "payment.export"
	methodDescription = `
Exports payment configuration from a workspace. Requires the 'payment.export'
permission in the target workspace.`
)

var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(methodKey),
	},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.Payment.Admin.Export(
			ctx.Context,
			data.WorkspaceID,
			padm.ExportRequest{Now: data.Now},
		)

		return Response{Package: value}, err
	},
}
