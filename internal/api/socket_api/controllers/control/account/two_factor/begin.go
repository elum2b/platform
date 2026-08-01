package twofactor

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type TwoFactorBeginRequest struct {
	Issuer string `json:"issuer" validate:"required,max=255"`
}

type TwoFactorBeginResponse struct {
	Secret string `json:"secret"`
	URI    string `json:"uri"`
}

func Begin(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		data := new(TwoFactorBeginRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		setup, err := services.Control.Admin.BeginTwoFactor(
			ctx,
			ctx.Peer.Identity().UserID,
			data.Issuer,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(ctx, event, TwoFactorBeginResponse{
			Secret: setup.Secret,
			URI:    setup.URI,
		})
	})
}
