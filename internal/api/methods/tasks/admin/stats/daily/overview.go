package daily

import (
	"time"

	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type OverviewRequest struct {
	WorkspaceID string    `json:"workspace_id" validate:"required,uuid"`
	From        time.Time `json:"from"         validate:"required"`
	Until       time.Time `json:"until"        validate:"required"`
}

type OverviewResponse struct {
	Stats []tadmin.DailyOverviewModel `json:"stats"`
}

var (
	dailyOverviewKey         = "tasks.stats.daily.overview"
	dailyOverviewDescription = `
Lists daily overview statistics for a workspace. Requires the
'tasks.stats.daily.overview' permission in the target workspace.`
)

// Overview exposes the daily overview statistics method.
var Overview = adapter.Method[OverviewRequest, OverviewResponse]{
	Key:         dailyOverviewKey,
	Description: dailyOverviewDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(dailyOverviewKey),
	},
	Handler: func(ctx *adapter.Context, data OverviewRequest) (OverviewResponse, error) {
		values, err := services.Tasks.Admin.ListDailyOverview(
			ctx.Context,
			data.WorkspaceID,
			data.From,
			data.Until,
		)

		return OverviewResponse{Stats: values}, err
	},
}
