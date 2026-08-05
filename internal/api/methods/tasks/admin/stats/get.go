package stats

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type GetResponse struct {
	Stats tadmin.StatsModel `json:"stats"`
}

var (
	getKey         = "tasks.stats.get"
	getDescription = `
Returns aggregate statistics for tasks in a workspace. Requires the
'tasks.stats.get' permission in the target workspace.`
)

// Get exposes the tasks aggregate statistics method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		value, err := services.Tasks.Admin.GetStats(
			ctx.Context,
			data.WorkspaceID,
		)

		return GetResponse{Stats: value}, err
	},
}
