package permission

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	RoleID      string `json:"role_id"      validate:"required,uuid"`
}

type ListResponse struct {
	MethodKeys []string `json:"method_keys"`
}

var (
	listKey         = "control.workspace.role.permission.list"
	listDescription = `
Lists method permissions assigned to a workspace role. Requires the
'control.workspace.role.permission.list' permission in the target workspace.`
)

// List returns permissions assigned to a workspace role.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.role.permission.list"),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		items, err := services.Control.Admin.ListWorkspaceRolePermissions(
			ctx.Context,
			data.WorkspaceID,
			data.RoleID,
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{MethodKeys: items}, nil
	},
}
