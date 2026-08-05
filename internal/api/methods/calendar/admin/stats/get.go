package stats

import (
	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CalendarID  string `json:"calendar_id"  validate:"required,max=255"`
}

type GetResponse struct {
	Stats caladmin.StatsModel `json:"stats"`
}

var (
	getKey         = "calendar.stats.get"
	getDescription = `
Returns aggregate statistics for a calendar. Requires the
'calendar.stats.get' permission in the target workspace.`
)

// Get exposes the calendar aggregate statistics method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		value, err := services.Calendar.Admin.GetStats(
			ctx.Context,
			data.WorkspaceID,
			data.CalendarID,
		)

		return GetResponse{Stats: value}, err
	},
}
