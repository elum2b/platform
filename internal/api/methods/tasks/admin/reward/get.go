package reward

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	TaskID      uint64 `json:"task_id"      validate:"required,min=1"`
	Key         string `json:"key"          validate:"required,max=255"`
}

type GetResponse struct {
	Reward tadmin.RewardModel `json:"reward"`
}

var (
	getKey         = "tasks.reward.get"
	getDescription = `
Returns a task reward. Requires the 'tasks.reward.get' permission in the
target workspace.`
)

// Get exposes the task reward retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		rwd, err := services.Tasks.Admin.GetReward(
			ctx.Context,
			data.WorkspaceID,
			data.TaskID,
			data.Key,
		)

		return GetResponse{Reward: rwd}, err
	},
}
