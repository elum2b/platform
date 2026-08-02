package auth

import (
	controlauth "github.com/elum2b/services/control/adapters"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type TelegramRequest struct {
	InitData    string `json:"init_data"    validate:"required"`
	InviteToken string `json:"invite_token"`
	BindToIP    bool   `json:"bind_to_ip"`
}

var (
	telegramKey         = "control.auth.telegram"
	telegramDescription = `
Authenticates an account through Telegram Mini Apps init data.`
)

// Telegram authenticates an account through Telegram Mini Apps init data.
var Telegram = adapter.Method[TelegramRequest, AuthResponse]{
	Key:         telegramKey,
	Description: telegramDescription,
	Transports:  adapter.HTTP,
	Method:      fiber.MethodPost,
	Handler: func(
		ctx *adapter.Context,
		data TelegramRequest,
	) (AuthResponse, error) {
		metadata := identityMetadata(ctx, data.InviteToken, data.BindToIP)

		identity, err := controlauth.TelegramWebApp(
			ctx.Context,
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
			return AuthResponse{}, err
		}

		return complete(ctx, identity)
	},
}
