package session

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type SessionRevokeRequest struct {
	SessionID string `json:"session_id" validate:"required,uuid"`
}

type SessionRevokeResponse struct {
	Affected int64 `json:"affected"`
}

func Revoke(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		data := new(SessionRevokeRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		affected, err := services.Control.Admin.RevokeSession(
			ctx,
			ctx.Peer.Identity().UserID,
			data.SessionID,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			SessionRevokeResponse{Affected: affected},
		)
	})
}
