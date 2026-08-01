package auth

import (
	controlauth "github.com/elum2b/services/control/adapters"
	serviceerrors "github.com/elum2b/services/errors"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
	httputils "github.com/elum2b/platform/internal/utils/http"
)

type TelegramRequest struct {
	InitData    string `json:"init_data"    validate:"required"`
	InviteToken string `json:"invite_token"`
	BindToIP    bool   `json:"bind_to_ip"`
}

func Telegram(path string, app fiber.Router) {
	app.Post(path, func(ctx fiber.Ctx) error {
		if !controlReady() {
			return httputils.Error(ctx, serviceerrors.ErrNotReady)
		}

		data := new(TelegramRequest)
		if !httputils.Decode(ctx, data) {
			return httputils.Error(ctx, serviceerrors.ErrInvalidFields)
		}

		metadata := identityMetadata(ctx, data.InviteToken, data.BindToIP)

		identity, err := controlauth.TelegramWebApp(
			ctx.Context(),
			controlauth.TelegramWebAppAuthParams{
				BotToken:    config.ControlAuthTelegramBotToken,
				InitData:    data.InitData,
				InviteToken: metadata.InviteToken,
				IP:          metadata.IP,
				UserAgent:   metadata.UserAgent,
				BindToIP:    metadata.BindToIP,
				ExpiresAt:   metadata.ExpiresAt,
				MaxAge:      config.ControlAuthTelegramMaxAge,
			},
		)
		if err != nil {
			return httputils.Error(ctx, err)
		}

		return complete(ctx, identity)
	})
}
