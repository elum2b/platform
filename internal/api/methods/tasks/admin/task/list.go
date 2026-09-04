package task

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id"        query:"workspace_id" validate:"required,uuid"`
	GroupKey    string `json:"group_key,omitempty" query:"group_key"`
	Limit       int32  `json:"limit,omitempty"     query:"limit"        validate:"omitempty,min=1,max=100"`
	Offset      int32  `json:"offset,omitempty"    query:"offset"       validate:"min=0"`
}

type ListResponse struct {
	Tasks []tadmin.TaskModel `json:"tasks"`
}

var (
	listKey         = "tasks.task.list"
	listDescription = `
Lists tasks in a workspace, optionally filtered by group. Requires the
'tasks.task.list' permission in the target workspace.`
)

// List exposes the task listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		tasks, err := services.Tasks.Admin.ListTasks(
			ctx.Context,
			data.WorkspaceID,
			data.GroupKey,
			data.Limit,
			data.Offset,
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Tasks: tasks}, nil
	},
}
