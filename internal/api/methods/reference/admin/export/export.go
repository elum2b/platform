package exportapi

import (
	"time"

	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type Request struct {
	WorkspaceID    string    `json:"workspace_id"               validate:"required,uuid"`
	Now            time.Time `json:"now,omitempty"`
	OnlyNotDeleted bool      `json:"only_not_deleted,omitempty"`
}

type Response struct {
	Package refadmin.ExportPackage `json:"package"`
}

var (
	methodKey         = "reference.export"
	methodDescription = `
Exports all reference items and their localizations from a workspace. Requires
the 'reference.export' permission in the target workspace.`
)

// Method exposes the reference export method.
var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(methodKey),
	},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.Reference.Admin.Export(
			ctx.Context,
			data.WorkspaceID,
			refadmin.ExportRequest{
				Now:            data.Now,
				OnlyNotDeleted: data.OnlyNotDeleted,
			},
		)

		return Response{Package: value}, err
	},
}
