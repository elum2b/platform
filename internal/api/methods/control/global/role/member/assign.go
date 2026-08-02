package member

import (
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type AssignRequest struct {
	AccountID string `json:"account_id" validate:"required,uuid"`
	RoleID    string `json:"role_id"    validate:"required,uuid"`
}

var (
	assignKey         = "control.global.role.member.assign"
	assignDescription = `
Assigns a global platform role to an account. Requires the
'control.global.role.member.assign' global permission.`
)

// Assign assigns a global role to an account.
var Assign = adapter.Method[AssignRequest, struct{}]{
	Key:         assignKey,
	Description: assignDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.role.member.assign"),
	},
	Handler: func(ctx *adapter.Context, data AssignRequest) (struct{}, error) {
		err := services.Control.Admin.AssignGlobalRole(
			ctx.Context,
			controladmin.SetRoleMemberParams{
				ActorID: ctx.AccountID, AccountID: data.AccountID,
				RoleID: data.RoleID,
			},
		)

		return struct{}{}, err
	},
}
