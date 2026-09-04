package importapi

import (
	"bytes"

	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	"github.com/elum2b/platform/internal/utils/archive"
)

type Request struct {
	WorkspaceID      string `json:"workspace_id"                query:"workspace_id"      validate:"required,uuid"`
	Archive          []byte `json:"archive"                     query:"archive"           validate:"required"`
	ConflictStrategy string `json:"conflict_strategy,omitempty" query:"conflict_strategy" validate:"omitempty,oneof=fail_on_conflict skip_existing update_existing"`
	IncludeMedia     bool   `json:"include_media,omitempty"     query:"include_media"`
}

type Response struct {
	Job refadmin.ArchiveJob `json:"job"`
}

var (
	methodKey         = "reference.import"
	methodDescription = `
Queues an archive import of reference items and localizations into a workspace. Requires the
'reference.import' permission in the target workspace.`
)

// Method exposes the reference import method.
var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(methodKey),
	},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.Reference.Admin.QueueArchiveImport(
			ctx.Context,
			refadmin.QueueArchiveImportParams{
				WorkspaceID:      data.WorkspaceID,
				FileName:         archive.FileName("reference"),
				IncludeMedia:     data.IncludeMedia,
				ConflictStrategy: data.ConflictStrategy,
				Archive:          bytes.NewReader(data.Archive),
			},
		)

		return Response{Job: value}, err
	},
}
