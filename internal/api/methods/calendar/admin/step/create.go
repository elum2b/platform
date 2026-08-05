package step

import (
	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CalendarID  string `json:"calendar_id"  validate:"required,max=255"`
	Position    uint32 `json:"position"     validate:"required,min=1"`
}

type CreateResponse struct {
	ID uint64 `json:"id"`
}

var (
	createKey         = "calendar.step.create"
	createDescription = `
Creates a new step in a calendar. Requires the 'calendar.step.create'
permission in the target workspace.`
)

// Create exposes the calendar step creation method.
var Create = adapter.Method[CreateRequest, CreateResponse]{
	Key:         createKey,
	Description: createDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(createKey),
	},
	Handler: func(ctx *adapter.Context, data CreateRequest) (CreateResponse, error) {
		id, err := services.Calendar.Admin.CreateStep(
			ctx.Context,
			caladmin.SaveStepParams{
				WorkspaceID: data.WorkspaceID,
				CalendarID:  data.CalendarID,
				Position:    data.Position,
			},
		)

		return CreateResponse{ID: id}, err
	},
}
