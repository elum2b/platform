package permission

import (
	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type ReplaceRequest struct {
	WorkspaceID string   `json:"workspace_id" validate:"required,uuid"`
	RoleID      string   `json:"role_id"      validate:"required,uuid"`
	MethodKeys  []string `json:"method_keys"  validate:"dive,required"`
}

func Replace(event string, socket etp.Router) {
	socket.Use(
		event,
		middleware.WorkspaceAccess("control.workspace.role.permission.replace"),
	)
	socket.On(event, func(ctx *etp.Context) error {
		data := new(ReplaceRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		if err := services.Control.Admin.ReplaceWorkspaceRolePermissions(
			ctx,
			admin.ReplaceRolePermissionsParams{
				ActorID:     ctx.Peer.Identity().UserID,
				WorkspaceID: data.WorkspaceID,
				RoleID:      data.RoleID,
				MethodKeys:  data.MethodKeys,
			},
		); err != nil {
			return err
		}

		return socketutils.Respond(ctx, event, struct{}{})
	})
}
