package twofactor

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DisableRequest struct {
	Code string `json:"code" validate:"required"`
}

type DisableResponse struct {
	Affected int64 `json:"affected"`
}

var (
	disableKey         = "control.account.two_factor.disable"
	disableDescription = `
Disables TOTP two-factor authentication for the account.`
)

// Disable disables two-factor authentication for the account.
var Disable = adapter.Method[DisableRequest, DisableResponse]{
	Key:         disableKey,
	Description: disableDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, data DisableRequest) (DisableResponse, error) {
		affected, err := services.Control.Admin.DisableTwoFactor(
			ctx.Context,
			ctx.AccountID,
			data.Code,
		)
		if err != nil {
			return DisableResponse{}, err
		}

		return DisableResponse{Affected: affected}, nil
	},
}
