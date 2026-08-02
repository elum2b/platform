package permission

import (
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ReplaceRequest struct {
	WorkspaceID string   `json:"workspace_id" validate:"required,uuid"`
	RoleID      string   `json:"role_id"      validate:"required,uuid"`
	MethodKeys  []string `json:"method_keys"  validate:"dive,required"`
}

var (
	replaceKey         = "control.workspace.role.permission.replace"
	replaceDescription = `
Replaces all method permissions assigned to a workspace role. Requires the
'control.workspace.role.permission.replace' permission in the target workspace.`
)

// Replace replaces permissions assigned to a workspace role.
var Replace = adapter.Method[ReplaceRequest, struct{}]{
	Key:         replaceKey,
	Description: replaceDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.role.permission.replace"),
	},
	Handler: func(ctx *adapter.Context, data ReplaceRequest) (struct{}, error) {
		err := services.Control.Admin.ReplaceWorkspaceRolePermissions(
			ctx.Context,
			controladmin.ReplaceRolePermissionsParams{
				ActorID: ctx.AccountID, WorkspaceID: data.WorkspaceID,
				RoleID: data.RoleID, MethodKeys: data.MethodKeys,
			},
		)

		return struct{}{}, err
	},
}
