package role

import (
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateRequest struct {
	RoleID      string `json:"role_id"               validate:"required,uuid"`
	Title       string `json:"title"                 validate:"required,max=255"`
	Description string `json:"description,omitempty" validate:"max=1000"`
	Position    int32  `json:"position,omitempty"`
}

type UpdateResponse struct {
	Affected int64 `json:"affected"`
}

var (
	updateKey         = "control.global.role.update"
	updateDescription = `
Updates a global platform role's title, description and position. Requires the
'control.global.role.update' global permission.`
)

// Update updates a global platform role.
var Update = adapter.Method[UpdateRequest, UpdateResponse]{
	Key:         updateKey,
	Description: updateDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.role.update"),
	},
	Handler: func(ctx *adapter.Context, data UpdateRequest) (UpdateResponse, error) {
		affected, err := services.Control.Admin.UpdateGlobalRole(
			ctx.Context,
			controladmin.UpdateRoleParams{
				ActorID: ctx.AccountID, ID: data.RoleID, Title: data.Title,
				Description: data.Description, Position: data.Position,
			},
		)
		if err != nil {
			return UpdateResponse{}, err
		}

		return UpdateResponse{Affected: affected}, nil
	},
}
