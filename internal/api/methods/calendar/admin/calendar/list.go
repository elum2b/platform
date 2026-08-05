package calendar

import (
	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id"     validate:"required,uuid"`
	Limit       int32  `json:"limit,omitempty"  validate:"omitempty,min=1,max=100"`
	Offset      int32  `json:"offset,omitempty" validate:"min=0"`
}

type ListResponse struct {
	Calendars []caladmin.CalendarModel `json:"calendars"`
}

var (
	listKey         = "calendar.list"
	listDescription = `
Lists calendars in a workspace. Requires the 'calendar.list' permission in the
target workspace.`
)

// List exposes the calendar listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		cals, err := services.Calendar.Admin.ListCalendars(
			ctx.Context,
			data.WorkspaceID,
			caladmin.Page{Limit: data.Limit, Offset: data.Offset},
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Calendars: cals}, nil
	},
}
