package export

import (
	"time"

	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	"github.com/elum2b/platform/internal/utils/archive"
)

type Request struct {
	WorkspaceID string    `json:"workspace_id"  query:"workspace_id" validate:"required,uuid"`
	Now         time.Time `json:"now,omitempty" query:"now"`
}

type Response struct {
	Job cpaadmin.ArchiveJob `json:"job"`
}

var (
	methodKey         = "cpa.export"
	methodDescription = `
Queues an archive export of all CPA offers and their configuration from a workspace. Requires the
'cpa.export' permission in the target workspace.`
)

// Method exposes the CPA export method.
var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(methodKey)},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.CPA.Admin.QueueArchiveExport(
			ctx.Context,
			cpaadmin.QueueArchiveExportParams{
				WorkspaceID:   data.WorkspaceID,
				FileName:      archive.FileName("cpa"),
				ExportRequest: cpaadmin.ExportRequest{Now: data.Now},
			},
		)

		return Response{Job: value}, err
	},
}
