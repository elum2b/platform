package importapi

import (
	"bytes"

	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	"github.com/elum2b/platform/internal/utils/archive"
)

type Request struct {
	WorkspaceID      string `json:"workspace_id"                validate:"required,uuid"`
	Archive          []byte `json:"archive"                     validate:"required"`
	ConflictStrategy string `json:"conflict_strategy,omitempty" validate:"omitempty,oneof=fail_on_conflict skip_existing update_existing"`
}

type Response struct {
	Job promoadmin.ArchiveJob `json:"job"`
}

var (
	methodKey         = "promo.import"
	methodDescription = `
Queues an archive import of promos and configuration into a workspace. Requires the 'promo.import'
permission in the target workspace.`
)

// Method exposes the promo import method.
var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(methodKey),
	},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.Promo.Admin.QueueArchiveImport(
			ctx.Context,
			promoadmin.QueueArchiveImportParams{
				WorkspaceID: data.WorkspaceID,
				FileName:    archive.FileName("promo"),
				ImportRequest: promoadmin.ImportRequest{
					ConflictStrategy: data.ConflictStrategy,
				},
				Archive: bytes.NewReader(data.Archive),
			},
		)

		return Response{Job: value}, err
	},
}
