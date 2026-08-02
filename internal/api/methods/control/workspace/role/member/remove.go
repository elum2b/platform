package member

import (
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type RemoveRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	AccountID   string `json:"account_id"   validate:"required,uuid"`
	RoleID      string `json:"role_id"      validate:"required,uuid"`
}

var (
	removeKey         = "control.workspace.role.member.remove"
	removeDescription = `
Removes a role from a workspace member. Requires the
'control.workspace.role.member.remove' permission in the target workspace.`
)

// Remove removes a workspace role from an account.
var Remove = adapter.Method[RemoveRequest, struct{}]{
	Key:         removeKey,
	Description: removeDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.role.member.remove"),
	},
	Handler: func(ctx *adapter.Context, data RemoveRequest) (struct{}, error) {
		err := services.Control.Admin.RemoveWorkspaceRole(
			ctx.Context,
			controladmin.SetRoleMemberParams{
				ActorID: ctx.AccountID, WorkspaceID: data.WorkspaceID,
				AccountID: data.AccountID, RoleID: data.RoleID,
			},
		)

		return struct{}{}, err
	},
}
