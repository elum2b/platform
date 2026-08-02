package role

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	RoleID string `json:"role_id" validate:"required,uuid"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "control.global.role.delete"
	deleteDescription = `
Deletes a global platform role. Requires the
'control.global.role.delete' global permission.`
)

// Delete deletes a global platform role.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.role.delete"),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Control.Admin.DeleteGlobalRole(
			ctx.Context,
			ctx.AccountID,
			data.RoleID,
		)
		if err != nil {
			return DeleteResponse{}, err
		}

		return DeleteResponse{Affected: affected}, nil
	},
}
