package importapi

import (
	"bytes"

	padm "github.com/elum2b/services/payment/service/admin"

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
	Job padm.ArchiveJob `json:"job"`
}

var (
	methodKey         = "payment.import"
	methodDescription = `
Queues an archive import of payment configuration into a workspace. Requires the 'payment.import'
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
		value, err := services.Payment.Admin.QueueArchiveImport(
			ctx.Context,
			padm.QueueArchiveImportParams{
				WorkspaceID: data.WorkspaceID,
				FileName:    archive.FileName("payment"),
				ImportRequest: padm.ImportRequest{
					ConflictStrategy: data.ConflictStrategy,
				},
				Archive: bytes.NewReader(data.Archive),
			},
		)

		return Response{Job: value}, err
	},
}
