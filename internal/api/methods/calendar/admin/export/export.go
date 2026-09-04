package exportapi

import (
	"time"

	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	"github.com/elum2b/platform/internal/utils/archive"
)

type Request struct {
	WorkspaceID string    `json:"workspace_id"  query:"workspace_id" validate:"required,uuid"`
	Now         time.Time `json:"now,omitempty" query:"now"`
}

type Response struct {
	Job caladmin.ArchiveJob `json:"job"`
}

var (
	methodKey         = "calendar.export"
	methodDescription = `
Queues an archive export of all calendars and their configuration from a workspace. Requires the
'calendar.export' permission in the target workspace.`
)

// Method exposes the calendar export method.
var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(methodKey),
	},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.Calendar.Admin.QueueArchiveExport(
			ctx.Context,
			caladmin.QueueArchiveExportParams{
				WorkspaceID:   data.WorkspaceID,
				FileName:      archive.FileName("calendar"),
				ExportRequest: caladmin.ExportRequest{Now: data.Now},
			},
		)

		return Response{Job: value}, err
	},
}
