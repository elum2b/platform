package identity

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type IdentityUnbindRequest struct {
	Provider string `json:"provider" validate:"required"`
}

type IdentityUnbindResponse struct {
	Affected int64 `json:"affected"`
}

func Unbind(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		data := new(IdentityUnbindRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		affected, err := services.Control.Admin.UnbindIdentity(
			ctx,
			ctx.Peer.Identity().UserID,
			data.Provider,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			IdentityUnbindResponse{Affected: affected},
		)
	})
}
