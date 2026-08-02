package member

import (
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type AssignRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	AccountID   string `json:"account_id"   validate:"required,uuid"`
	RoleID      string `json:"role_id"      validate:"required,uuid"`
}

var (
	assignKey         = "control.workspace.role.member.assign"
	assignDescription = `
Assigns a role to a workspace member. Requires the
'control.workspace.role.member.assign' permission in the target workspace.`
)

// Assign assigns a workspace role to an account.
var Assign = adapter.Method[AssignRequest, struct{}]{
	Key:         assignKey,
	Description: assignDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.role.member.assign"),
	},
	Handler: func(ctx *adapter.Context, data AssignRequest) (struct{}, error) {
		err := services.Control.Admin.AssignWorkspaceRole(
			ctx.Context,
			controladmin.SetRoleMemberParams{
				ActorID: ctx.AccountID, WorkspaceID: data.WorkspaceID,
				AccountID: data.AccountID, RoleID: data.RoleID,
			},
		)

		return struct{}{}, err
	},
}
