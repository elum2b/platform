package exportapi

import (
	"time"

	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	"github.com/elum2b/platform/internal/utils/archive"
)

type Request struct {
	WorkspaceID string    `json:"workspace_id"  query:"workspace_id" validate:"required,uuid"`
	Now         time.Time `json:"now,omitempty" query:"now"`
}

type Response struct {
	Job promoadmin.ArchiveJob `json:"job"`
}

var (
	methodKey         = "promo.export"
	methodDescription = `
Queues an archive export of all promos and their configuration from a workspace. Requires the
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
		value, err := services.Promo.Admin.QueueArchiveExport(
			ctx.Context,
			promoadmin.QueueArchiveExportParams{
				WorkspaceID:   data.WorkspaceID,
				FileName:      archive.FileName("promo"),
				ExportRequest: promoadmin.ExportRequest{Now: data.Now},
			},
		)

		return Response{Job: value}, err
	},
}
