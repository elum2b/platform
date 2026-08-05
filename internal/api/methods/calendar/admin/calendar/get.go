package calendar

import (
	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          string `json:"id"           validate:"required,max=255"`
}

type GetResponse struct {
	Calendar caladmin.CalendarModel `json:"calendar"`
}

var (
	getKey         = "calendar.get"
	getDescription = `
Returns a calendar with its localizations, steps and rewards. Requires the
'calendar.get' permission in the target workspace.`
)

// Get exposes the calendar retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		cal, err := services.Calendar.Admin.GetCalendar(
			ctx.Context,
			data.WorkspaceID,
			data.ID,
		)

		return GetResponse{Calendar: cal}, err
	},
}
