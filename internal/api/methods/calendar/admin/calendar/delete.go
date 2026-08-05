package calendar

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          string `json:"id"           validate:"required,max=255"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "calendar.delete"
	deleteDescription = `
Deletes a calendar and its dependent data. Requires the 'calendar.delete'
permission in the target workspace.`
)

// Delete exposes the calendar deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(deleteKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Calendar.Admin.DeleteCalendar(
			ctx.Context,
			data.WorkspaceID,
			data.ID,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
