package exportapi

import (
	"time"

	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type Request struct {
	WorkspaceID string    `json:"workspace_id"  validate:"required,uuid"`
	Now         time.Time `json:"now,omitempty"`
}

type Response struct {
	Package promoadmin.ExportPackage `json:"package"`
}

var (
	methodKey         = "promo.export"
	methodDescription = `
Exports all promos and their configuration from a workspace. Requires the
'promo.export' permission in the target workspace.`
)

// Method exposes the promo export method.
var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(methodKey),
	},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.Promo.Admin.Export(
			ctx.Context,
			data.WorkspaceID,
			promoadmin.ExportRequest{Now: data.Now},
		)

		return Response{Package: value}, err
	},
}
