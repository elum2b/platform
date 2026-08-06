package group

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type ListResponse struct {
	Groups []tadmin.GroupModel `json:"groups"`
}

var (
	listKey         = "tasks.group.list"
	listDescription = `
Lists task groups in a workspace. Requires the 'tasks.group.list' permission
in the target workspace.`
)

// List exposes the task group listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		groups, err := services.Tasks.Admin.ListGroups(
			ctx.Context,
			data.WorkspaceID,
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Groups: groups}, nil
	},
}
