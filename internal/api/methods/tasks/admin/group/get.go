package group

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Key         string `json:"key"          validate:"required,max=255"`
}

type GetResponse struct {
	Group tadmin.GroupModel `json:"group"`
}

var (
	getKey         = "tasks.group.get"
	getDescription = `
Returns a task group. Requires the 'tasks.group.get' permission in the target
workspace.`
)

// Get exposes the task group retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		group, err := services.Tasks.Admin.GetGroup(
			ctx.Context,
			data.WorkspaceID,
			data.Key,
		)

		return GetResponse{Group: group}, err
	},
}
