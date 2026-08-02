package role

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	RoleID      string `json:"role_id"      validate:"required,uuid"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "control.workspace.role.delete"
	deleteDescription = `
Deletes a workspace role. Requires the
'control.workspace.role.delete' permission in the target workspace.`
)

// Delete deletes a workspace role.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.role.delete"),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Control.Admin.DeleteWorkspaceRole(
			ctx.Context,
			ctx.AccountID,
			data.WorkspaceID,
			data.RoleID,
		)
		if err != nil {
			return DeleteResponse{}, err
		}

		return DeleteResponse{Affected: affected}, nil
	},
}
