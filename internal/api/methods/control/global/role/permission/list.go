package permission

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	RoleID string `json:"role_id" validate:"required,uuid"`
}

type ListResponse struct {
	MethodKeys []string `json:"method_keys"`
}

var (
	listKey         = "control.global.role.permission.list"
	listDescription = `
Lists method permissions assigned to a global platform role. Requires the
'control.global.role.permission.list' global permission.`
)

// List returns permissions assigned to a global role.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.role.permission.list"),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		items, err := services.Control.Admin.ListGlobalRolePermissions(
			ctx.Context,
			data.RoleID,
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{MethodKeys: items}, nil
	},
}
