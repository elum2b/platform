package importapi

import (
	"bytes"

	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	"github.com/elum2b/platform/internal/utils/archive"
)

type Request struct {
	WorkspaceID      string `json:"workspace_id"                query:"workspace_id"      validate:"required,uuid"`
	Archive          []byte `json:"archive"                     query:"archive"           validate:"required"`
	ConflictStrategy string `json:"conflict_strategy,omitempty" query:"conflict_strategy" validate:"omitempty,oneof=fail_on_conflict skip_existing update_existing"`
}

type Response struct {
	Job caladmin.ArchiveJob `json:"job"`
}

var (
	methodKey         = "calendar.import"
	methodDescription = `
Queues an archive import of calendars and configuration into a workspace. Requires the
'calendar.import' permission in the target workspace.`
)

// Method exposes the calendar import method.
var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(methodKey),
	},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.Calendar.Admin.QueueArchiveImport(
			ctx.Context,
			caladmin.QueueArchiveImportParams{
				WorkspaceID: data.WorkspaceID,
				FileName:    archive.FileName("calendar"),
				ImportRequest: caladmin.ImportRequest{
					ConflictStrategy: data.ConflictStrategy,
				},
				Archive: bytes.NewReader(data.Archive),
			},
		)

		return Response{Job: value}, err
	},
}
