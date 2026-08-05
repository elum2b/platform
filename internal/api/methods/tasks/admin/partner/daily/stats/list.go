package stats

import (
	"time"

	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string    `json:"workspace_id" validate:"required,uuid"`
	Provider    string    `json:"provider"     validate:"required"`
	GroupKey    string    `json:"group_key"    validate:"required,max=255"`
	From        time.Time `json:"from"         validate:"required"`
	Until       time.Time `json:"until"        validate:"required"`
}

type ListResponse struct {
	Stats []tadmin.PartnerDailyStatsModel `json:"stats"`
}

var (
	dailyStatsListKey         = "tasks.partner.daily_stats.list"
	dailyStatsListDescription = `
Lists daily partner statistics. Requires the
'tasks.partner.daily_stats.list' permission in the target workspace.`
)

// List exposes the partner daily stats listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         dailyStatsListKey,
	Description: dailyStatsListDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(dailyStatsListKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.Tasks.Admin.ListPartnerDailyStats(
			ctx.Context,
			data.WorkspaceID,
			data.Provider,
			data.GroupKey,
			data.From,
			data.Until,
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Stats: values}, nil
	},
}
