package member

import (
	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type RemoveRequest struct {
	AccountID string `json:"account_id" validate:"required,uuid"`
	RoleID    string `json:"role_id"    validate:"required,uuid"`
}

func Remove(event string, socket etp.Router) {
	socket.Use(
		event,
		middleware.GlobalAccess("control.global.role.member.remove"),
	)
	socket.On(event, func(ctx *etp.Context) error {
		data := new(RemoveRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		if err := services.Control.Admin.RemoveGlobalRole(
			ctx,
			admin.SetRoleMemberParams{
				ActorID:   ctx.Peer.Identity().UserID,
				AccountID: data.AccountID,
				RoleID:    data.RoleID,
			},
		); err != nil {
			return err
		}

		return socketutils.Respond(ctx, event, struct{}{})
	})
}
