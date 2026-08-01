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
	RoleID     string   `json:"role_id"     validate:"required,uuid"`
	MethodKeys []string `json:"method_keys" validate:"dive,required"`
}

func Replace(event string, socket etp.Router) {
	socket.Use(
		event,
		middleware.GlobalAccess("control.global.role.permission.replace"),
	)
	socket.On(event, func(ctx *etp.Context) error {
		data := new(ReplaceRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		if err := services.Control.Admin.ReplaceGlobalRolePermissions(
			ctx,
			admin.ReplaceRolePermissionsParams{
				ActorID:    ctx.Peer.Identity().UserID,
				RoleID:     data.RoleID,
				MethodKeys: data.MethodKeys,
			},
		); err != nil {
			return err
		}

		return socketutils.Respond(ctx, event, struct{}{})
	})
}
