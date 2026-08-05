package step

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CalendarID  string `json:"calendar_id"  validate:"required,max=255"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "calendar.step.delete"
	deleteDescription = `
Deletes a step from a calendar. Requires the 'calendar.step.delete' permission
in the target workspace.`
)

// Delete exposes the calendar step deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(deleteKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Calendar.Admin.DeleteStep(
			ctx.Context,
			data.WorkspaceID,
			data.CalendarID,
			data.ID,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
