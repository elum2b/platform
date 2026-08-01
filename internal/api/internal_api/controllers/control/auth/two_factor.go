package auth

import (
	serviceerrors "github.com/elum2b/services/errors"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/services"
	"github.com/elum2b/platform/internal/utils/cookies"
	httputils "github.com/elum2b/platform/internal/utils/http"
)

type TwoFactorRequest struct {
	Code string `json:"code" validate:"required"`
}

func TwoFactor(path string, app fiber.Router) {
	app.Post(path, func(ctx fiber.Ctx) error {
		if !controlReady() {
			return httputils.Error(ctx, serviceerrors.ErrNotReady)
		}

		data := new(TwoFactorRequest)
		if !httputils.Decode(ctx, data) {
			return httputils.Error(ctx, serviceerrors.ErrInvalidFields)
		}

		challenge := cookies.Get(ctx, config.ControlAuthTwoFactorCookieName)
		if challenge == "" {
			return httputils.Error(ctx, serviceerrors.ErrUnauthorized)
		}

		result, err := services.Control.Admin.CompleteTwoFactor(
			ctx.Context(),
			challenge,
			data.Code,
			ctx.IP(),
		)
		if err != nil {
			return httputils.Error(ctx, err)
		}

		return respondAuth(ctx, result)
	})
}
