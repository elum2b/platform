package role

import (
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateRequest struct {
	WorkspaceID string `json:"workspace_id"          validate:"required,uuid"`
	RoleID      string `json:"role_id"               validate:"required,uuid"`
	Title       string `json:"title"                 validate:"required,max=255"`
	Description string `json:"description,omitempty" validate:"max=1000"`
	Position    int32  `json:"position,omitempty"`
}

type UpdateResponse struct {
	Affected int64 `json:"affected"`
}

var (
	updateKey         = "control.workspace.role.update"
	updateDescription = `
Updates a workspace role's title, description and position. Requires the
'control.workspace.role.update' permission in the target workspace.`
)

// Update updates a workspace role.
var Update = adapter.Method[UpdateRequest, UpdateResponse]{
	Key:         updateKey,
	Description: updateDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.role.update"),
	},
	Handler: func(ctx *adapter.Context, data UpdateRequest) (UpdateResponse, error) {
		affected, err := services.Control.Admin.UpdateWorkspaceRole(
			ctx.Context,
			controladmin.UpdateRoleParams{
				ActorID: ctx.AccountID, WorkspaceID: data.WorkspaceID,
				ID: data.RoleID, Title: data.Title,
				Description: data.Description, Position: data.Position,
			},
		)
		if err != nil {
			return UpdateResponse{}, err
		}

		return UpdateResponse{Affected: affected}, nil
	},
}
