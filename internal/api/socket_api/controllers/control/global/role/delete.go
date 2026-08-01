package role

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type DeleteRequest struct {
	RoleID string `json:"role_id" validate:"required,uuid"`
}
type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

func Delete(event string, socket etp.Router) {
	socket.Use(event, middleware.GlobalAccess("control.global.role.delete"))
	socket.On(event, func(ctx *etp.Context) error {
		data := new(DeleteRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		affected, err := services.Control.Admin.DeleteGlobalRole(
			ctx,
			ctx.Peer.Identity().UserID,
			data.RoleID,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			DeleteResponse{Affected: affected},
		)
	})
}
