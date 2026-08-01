package identity

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
	authutils "github.com/elum2b/platform/internal/utils/auth"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type IdentityBindRequest authutils.IdentityRequest

type IdentityBindResponse struct {
	Bound bool `json:"bound"`
}

func Bind(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		data := new(IdentityBindRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		identity, err := authutils.ResolveIdentity(
			ctx,
			authutils.IdentityRequest(*data),
		)
		if err != nil {
			return err
		}

		if err := services.Control.Admin.BindIdentity(
			ctx,
			ctx.Peer.Identity().UserID,
			identity,
		); err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			IdentityBindResponse{Bound: true},
		)
	})
}
