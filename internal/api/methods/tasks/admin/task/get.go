package task

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
}

type GetResponse struct {
	Task tadmin.TaskModel `json:"task"`
}

var (
	getKey         = "tasks.task.get"
	getDescription = `
Returns a task. Requires the 'tasks.task.get' permission in the target
workspace.`
)

// Get exposes the task retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		task, err := services.Tasks.Admin.GetTask(
			ctx.Context,
			data.WorkspaceID,
			data.ID,
		)

		return GetResponse{Task: task}, err
	},
}
