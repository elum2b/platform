package member

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type RemoveRequest struct {
	AccountID string `json:"account_id" validate:"required,uuid"`
}
type RemoveResponse struct {
	Affected int64 `json:"affected"`
}

func Remove(event string, socket etp.Router) {
	socket.Use(event, middleware.GlobalAccess("control.global.member.remove"))
	socket.On(event, func(ctx *etp.Context) error {
		data := new(RemoveRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		affected, err := services.Control.Admin.RemovePlatformMember(
			ctx,
			ctx.Peer.Identity().UserID,
			data.AccountID,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			RemoveResponse{Affected: affected},
		)
	})
}
