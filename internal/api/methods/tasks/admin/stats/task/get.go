package task

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type TaskGetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	TaskID      uint64 `json:"task_id"      validate:"required,min=1"`
}

type TaskGetResponse struct {
	Stats tadmin.TaskStatsModel `json:"stats"`
}

var (
	taskGetKey         = "tasks.stats.task.get"
	taskGetDescription = `
Returns statistics for a single task. Requires the 'tasks.stats.task.get'
permission in the target workspace.`
)

// TaskGet exposes the per-task statistics method.
var TaskGet = adapter.Method[TaskGetRequest, TaskGetResponse]{
	Key:         taskGetKey,
	Description: taskGetDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(taskGetKey),
	},
	Handler: func(ctx *adapter.Context, data TaskGetRequest) (TaskGetResponse, error) {
		value, err := services.Tasks.Admin.GetTaskStats(
			ctx.Context,
			data.WorkspaceID,
			data.TaskID,
		)

		return TaskGetResponse{Stats: value}, err
	},
}
