package twofactor

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type TwoFactorDisableRequest struct {
	Code string `json:"code" validate:"required"`
}

type TwoFactorDisableResponse struct {
	Affected int64 `json:"affected"`
}

func Disable(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		data := new(TwoFactorDisableRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		affected, err := services.Control.Admin.DisableTwoFactor(
			ctx,
			ctx.Peer.Identity().UserID,
			data.Code,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			TwoFactorDisableResponse{Affected: affected},
		)
	})
}
