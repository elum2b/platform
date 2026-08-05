package step

import (
	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CalendarID  string `json:"calendar_id"  validate:"required,max=255"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
	Position    uint32 `json:"position"     validate:"required,min=1"`
}

type UpdateResponse struct {
	Affected int64 `json:"affected"`
}

var (
	updateKey         = "calendar.step.update"
	updateDescription = `
Updates the position of a calendar step. Requires the
'calendar.step.update' permission in the target workspace.`
)

// Update exposes the calendar step update method.
var Update = adapter.Method[UpdateRequest, UpdateResponse]{
	Key:         updateKey,
	Description: updateDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(updateKey),
	},
	Handler: func(ctx *adapter.Context, data UpdateRequest) (UpdateResponse, error) {
		affected, err := services.Calendar.Admin.UpdateStep(
			ctx.Context,
			caladmin.SaveStepParams{
				WorkspaceID: data.WorkspaceID,
				CalendarID:  data.CalendarID,
				ID:          data.ID,
				Position:    data.Position,
			},
		)

		return UpdateResponse{Affected: affected}, err
	},
}
