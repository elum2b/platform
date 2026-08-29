package exportapi

import (
	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	"github.com/elum2b/platform/internal/utils/archive"
)

type Request struct {
	WorkspaceID  string `json:"workspace_id"             validate:"required,uuid"`
	IncludeMedia bool   `json:"include_media,omitempty"`
}

type Response struct {
	Job refadmin.ArchiveJob `json:"job"`
}

var (
	methodKey         = "reference.export"
	methodDescription = `
Queues an archive export of reference items and their localizations from a workspace. Requires
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
		value, err := services.Reference.Admin.QueueArchiveExport(
			ctx.Context,
			refadmin.QueueArchiveExportParams{
				WorkspaceID:  data.WorkspaceID,
				FileName:     archive.FileName("reference"),
				IncludeMedia: data.IncludeMedia,
			},
		)

		return Response{Job: value}, err
	},
}
