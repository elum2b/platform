package permission

import (
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ReplaceRequest struct {
	RoleID     string   `json:"role_id"     validate:"required,uuid"`
	MethodKeys []string `json:"method_keys" validate:"dive,required"`
}

var (
	replaceKey         = "control.global.role.permission.replace"
	replaceDescription = `
Replaces all method permissions assigned to a global platform role. Requires the
'control.global.role.permission.replace' global permission.`
)

// Replace replaces permissions assigned to a global role.
var Replace = adapter.Method[ReplaceRequest, struct{}]{
	Key:         replaceKey,
	Description: replaceDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.role.permission.replace"),
	},
	Handler: func(ctx *adapter.Context, data ReplaceRequest) (struct{}, error) {
		err := services.Control.Admin.ReplaceGlobalRolePermissions(
			ctx.Context,
			controladmin.ReplaceRolePermissionsParams{
				ActorID: ctx.AccountID, RoleID: data.RoleID,
				MethodKeys: data.MethodKeys,
			},
		)

		return struct{}{}, err
	},
}
