package twofactor

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ConfirmRequest struct {
	Code string `json:"code" validate:"required"`
}

type ConfirmResponse struct {
	BackupCodes []string `json:"backup_codes"`
}

var (
	confirmKey         = "control.account.two_factor.confirm"
	confirmDescription = `
Confirms TOTP setup and returns backup codes.`
)

// Confirm enables two-factor authentication using a TOTP code.
var Confirm = adapter.Method[ConfirmRequest, ConfirmResponse]{
	Key:         confirmKey,
	Description: confirmDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, data ConfirmRequest) (ConfirmResponse, error) {
		codes, err := services.Control.Admin.ConfirmTwoFactor(
			ctx.Context,
			ctx.AccountID,
			data.Code,
		)
		if err != nil {
			return ConfirmResponse{}, err
		}

		return ConfirmResponse{BackupCodes: codes}, nil
	},
}
