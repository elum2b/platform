package member

import (
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type RemoveRequest struct {
	AccountID string `json:"account_id" validate:"required,uuid"`
	RoleID    string `json:"role_id"    validate:"required,uuid"`
}

var (
	removeKey         = "control.global.role.member.remove"
	removeDescription = `
Removes a global platform role from an account. Requires the
'control.global.role.member.remove' global permission.`
)

// Remove removes a global role from an account.
var Remove = adapter.Method[RemoveRequest, struct{}]{
	Key:         removeKey,
	Description: removeDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.role.member.remove"),
	},
	Handler: func(ctx *adapter.Context, data RemoveRequest) (struct{}, error) {
		err := services.Control.Admin.RemoveGlobalRole(
			ctx.Context,
			controladmin.SetRoleMemberParams{
				ActorID: ctx.AccountID, AccountID: data.AccountID,
				RoleID: data.RoleID,
			},
		)

		return struct{}{}, err
	},
}
