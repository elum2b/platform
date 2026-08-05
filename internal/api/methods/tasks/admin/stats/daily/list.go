package daily

import (
	"time"

	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string    `json:"workspace_id" validate:"required,uuid"`
	TaskID      uint64    `json:"task_id"      validate:"required,min=1"`
	From        time.Time `json:"from"         validate:"required"`
	Until       time.Time `json:"until"        validate:"required"`
}

type ListResponse struct {
	Stats []tadmin.DailyStatsModel `json:"stats"`
}

var (
	dailyListKey         = "tasks.stats.daily.list"
	dailyListDescription = `
Lists daily statistics for a task. Requires the 'tasks.stats.daily.list'
permission in the target workspace.`
)

// List exposes the daily task statistics method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         dailyListKey,
	Description: dailyListDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(dailyListKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.Tasks.Admin.ListDailyStats(
			ctx.Context,
			data.WorkspaceID,
			data.TaskID,
			data.From,
			data.Until,
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Stats: values}, nil
	},
}
