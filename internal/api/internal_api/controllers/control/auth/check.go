package auth

import (
	serviceerrors "github.com/elum2b/services/errors"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/services"
	"github.com/elum2b/platform/internal/utils/cookies"
	httputils "github.com/elum2b/platform/internal/utils/http"
)

func Check(path string, app fiber.Router) {
	app.Get(path, func(ctx fiber.Ctx) error {
		if !controlReady() {
			return httputils.Error(ctx, serviceerrors.ErrNotReady)
		}

		token := cookies.Get(ctx, config.ControlAuthCookieName)
		if token == "" {
			return httputils.Respond(ctx, AuthResponse{Authenticated: false})
		}

		session, err := services.Control.Admin.ValidateSession(
			ctx.Context(),
			token,
			ctx.IP(),
		)
		if err != nil {
			if invalidSessionError(err) {
				cookies.Clear(ctx, config.ControlAuthCookieName)

				return httputils.Respond(
					ctx,
					AuthResponse{Authenticated: false},
				)
			}

			return httputils.Error(ctx, err)
		}

		account, err := services.Control.Admin.GetAccount(
			ctx.Context(),
			session.AccountID,
		)
		if err != nil {
			return httputils.Error(ctx, err)
		}

		return httputils.Respond(ctx, AuthResponse{
			Authenticated:    true,
			Account:          accountResponse(account),
			SessionID:        session.ID,
			SessionExpiresAt: session.ExpiresAt,
		})
	})
}

func invalidSessionError(err error) bool {
	switch serviceerrors.CodeOf(err) {
	case serviceerrors.CodeUnauthorized,
		serviceerrors.CodeForbidden,
		serviceerrors.CodeNotFound,
		serviceerrors.CodeFailedPrecondition:
		return true
	default:
		return false
	}
}
