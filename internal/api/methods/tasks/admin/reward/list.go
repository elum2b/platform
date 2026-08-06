package reward

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	TaskID      uint64 `json:"task_id"      validate:"required,min=1"`
}

type ListResponse struct {
	Rewards []tadmin.RewardModel `json:"rewards"`
}

var (
	listKey         = "tasks.reward.list"
	listDescription = `
Lists rewards for a task. Requires the 'tasks.reward.list' permission in the
target workspace.`
)

// List exposes the task reward listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		rwds, err := services.Tasks.Admin.ListRewards(
			ctx.Context,
			data.WorkspaceID,
			data.TaskID,
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Rewards: rwds}, nil
	},
}
