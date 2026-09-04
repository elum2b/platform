package exportapi

import (
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	"github.com/elum2b/platform/internal/utils/archive"
)

type Request struct {
	WorkspaceID string    `json:"workspace_id"  query:"workspace_id" validate:"required,uuid"`
	Now         time.Time `json:"now,omitempty" query:"now"`
}

type Response struct {
	Job padm.ArchiveJob `json:"job"`
}

var (
	methodKey         = "payment.export"
	methodDescription = `
Queues an archive export of payment configuration from a workspace. Requires the 'payment.export'
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
		value, err := services.Payment.Admin.QueueArchiveExport(
			ctx.Context,
			padm.QueueArchiveExportParams{
				WorkspaceID:   data.WorkspaceID,
				FileName:      archive.FileName("payment"),
				ExportRequest: padm.ExportRequest{Now: data.Now},
			},
		)

		return Response{Job: value}, err
	},
}
