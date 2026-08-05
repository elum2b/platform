package stats

import (
	"time"

	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DailyListRequest struct {
	WorkspaceID string    `json:"workspace_id" validate:"required,uuid"`
	CalendarID  string    `json:"calendar_id"  validate:"required,max=255"`
	From        time.Time `json:"from"         validate:"required"`
	Until       time.Time `json:"until"        validate:"required"`
}

type DailyListResponse struct {
	Stats []caladmin.DailyStatsModel `json:"stats"`
}

var (
	dailyListKey         = "calendar.stats.daily.list"
	dailyListDescription = `
Lists daily statistics for a calendar and time range. Requires the
'calendar.stats.daily.list' permission in the target workspace.`
)

// DailyList exposes the calendar daily statistics listing method.
var DailyList = adapter.Method[DailyListRequest, DailyListResponse]{
	Key:         dailyListKey,
	Description: dailyListDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(dailyListKey),
	},
	Handler: func(ctx *adapter.Context, data DailyListRequest) (DailyListResponse, error) {
		values, err := services.Calendar.Admin.ListDailyStats(
			ctx.Context,
			data.WorkspaceID,
			data.CalendarID,
			data.From,
			data.Until,
		)

		return DailyListResponse{Stats: values}, err
	},
}
