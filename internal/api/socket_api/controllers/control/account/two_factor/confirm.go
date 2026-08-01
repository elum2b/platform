package twofactor

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type TwoFactorConfirmRequest struct {
	Code string `json:"code" validate:"required"`
}

type TwoFactorConfirmResponse struct {
	BackupCodes []string `json:"backup_codes"`
}

func Confirm(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		data := new(TwoFactorConfirmRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		backupCodes, err := services.Control.Admin.ConfirmTwoFactor(
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
			TwoFactorConfirmResponse{BackupCodes: backupCodes},
		)
	})
}
