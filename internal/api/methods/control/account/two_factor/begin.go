package twofactor

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type BeginRequest struct {
	Issuer string `json:"issuer" validate:"required,max=255"`
}

type BeginResponse struct {
	Secret string `json:"secret"`
	URI    string `json:"uri"`
}

var (
	beginKey         = "control.account.two_factor.begin"
	beginDescription = `
Starts TOTP two-factor setup for the authenticated account.`
)

// Begin starts two-factor setup for the authenticated account.
var Begin = adapter.Method[BeginRequest, BeginResponse]{
	Key:         beginKey,
	Description: beginDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, data BeginRequest) (BeginResponse, error) {
		setup, err := services.Control.Admin.BeginTwoFactor(
			ctx.Context,
			ctx.AccountID,
			data.Issuer,
		)
		if err != nil {
			return BeginResponse{}, err
		}

		return BeginResponse{Secret: setup.Secret, URI: setup.URI}, nil
	},
}
