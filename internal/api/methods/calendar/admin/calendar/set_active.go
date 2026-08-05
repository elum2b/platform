package calendar

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type SetActiveRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          string `json:"id"           validate:"required,max=255"`
	Active      bool   `json:"active"`
}

type SetActiveResponse struct {
	Affected int64 `json:"affected"`
}

var (
	setActiveKey         = "calendar.set_active"
	setActiveDescription = `
Activates or deactivates a calendar. Requires the 'calendar.set_active'
permission in the target workspace.`
)

// SetActive exposes the calendar activation toggle method.
var SetActive = adapter.Method[SetActiveRequest, SetActiveResponse]{
	Key:         setActiveKey,
	Description: setActiveDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(setActiveKey),
	},
	Handler: func(ctx *adapter.Context, data SetActiveRequest) (SetActiveResponse, error) {
		affected, err := services.Calendar.Admin.SetCalendarActive(
			ctx.Context,
			data.WorkspaceID,
			data.ID,
			data.Active,
		)

		return SetActiveResponse{Affected: affected}, err
	},
}
