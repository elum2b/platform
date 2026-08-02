package auth

import (
	serviceerrors "github.com/elum2b/services/errors"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	"github.com/elum2b/platform/internal/utils/cookies"
)

type TwoFactorRequest struct {
	Code string `json:"code" validate:"required"`
}

var (
	twoFactorKey         = "control.auth.twoFactor"
	twoFactorDescription = `
Completes authentication using a two-factor code.`
)

// TwoFactor completes an HTTP authentication that requires a TOTP code.
var TwoFactor = adapter.Method[TwoFactorRequest, AuthResponse]{
	Key:         twoFactorKey,
	Description: twoFactorDescription,
	Transports:  adapter.HTTP,
	Method:      fiber.MethodPost,
	Handler: func(
		ctx *adapter.Context,
		data TwoFactorRequest,
	) (AuthResponse, error) {
		challenge := cookies.Get(
			ctx.HTTP,
			config.ControlAuthTwoFactorCookieName,
		)
		if challenge == "" {
			return AuthResponse{}, serviceerrors.ErrUnauthorized
		}

		result, err := services.Control.Admin.CompleteTwoFactor(
			ctx.Context,
			challenge,
			data.Code,
			ctx.HTTP.IP(),
		)
		if err != nil {
			return AuthResponse{}, err
		}

		return authResponse(ctx, result), nil
	},
}
