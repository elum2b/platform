package auth

import (
	serviceerrors "github.com/elum2b/services/errors"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	"github.com/elum2b/platform/internal/utils/cookies"
)

var (
	checkKey         = "control.auth.check"
	checkDescription = `
Gets the current authentication state from the session cookie.`
)

// Check exposes the current HTTP cookie authentication state.
var Check = adapter.Method[struct{}, AuthResponse]{
	Key:         checkKey,
	Description: checkDescription,
	Transports:  adapter.HTTP,
	Method:      fiber.MethodGet,
	Handler: func(ctx *adapter.Context, _ struct{}) (AuthResponse, error) {
		token := cookies.Get(ctx.HTTP, config.ControlAuthCookieName)
		if token == "" {
			return AuthResponse{Authenticated: false}, nil
		}

		session, err := services.Control.Admin.ValidateSession(
			ctx.Context,
			token,
			ctx.HTTP.IP(),
		)
		if err != nil {
			switch serviceerrors.CodeOf(err) {
			case serviceerrors.CodeUnauthorized,
				serviceerrors.CodeForbidden,
				serviceerrors.CodeNotFound,
				serviceerrors.CodeFailedPrecondition:
				cookies.Clear(ctx.HTTP, config.ControlAuthCookieName)

				return AuthResponse{Authenticated: false}, nil

			default:
				return AuthResponse{}, err
			}
		}

		account, err := services.Control.Admin.GetAccount(
			ctx.Context,
			session.AccountID,
		)
		if err != nil {
			return AuthResponse{}, err
		}

		return AuthResponse{
			Authenticated:    true,
			Account:          accountResponse(account),
			SessionID:        session.ID,
			SessionExpiresAt: session.ExpiresAt,
		}, nil
	},
}
