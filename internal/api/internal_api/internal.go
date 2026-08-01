package internalapi

import (
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/api/internal_api/controllers/control/auth"
)

func Init(app fiber.Router) {
	internal := app.Group("/internal")

	control := internal.Group("/control")

	auth.Check("auth.check", control)

	auth.VK("auth.vkid", control)

	auth.Telegram("auth.telegram", control)

	auth.GitHub("auth.github", control)

	auth.GitLab("auth.gitlab", control)

	auth.Google("auth.google", control)

	auth.Yandex("auth.yandex", control)

	auth.TON("auth.ton", control)
	auth.TONChallenge("auth.ton.challenge", control)

	auth.TwoFactor("auth.twoFactor", control)
}
