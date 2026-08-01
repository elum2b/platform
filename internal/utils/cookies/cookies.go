package cookies

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
)

func Get(ctx fiber.Ctx, name string) string {
	return ctx.Cookies(name)
}

func Set(ctx fiber.Ctx, name, value string, expiresAt time.Time) {
	ctx.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Domain:   config.ControlAuthCookieDomain,
		Expires:  expiresAt,
		Secure:   true,
		HTTPOnly: true,
		SameSite: config.ControlAuthCookieSameSite,
	})
}

func Clear(ctx fiber.Ctx, name string) {
	ctx.Cookie(&fiber.Cookie{
		Name:     name,
		Path:     "/",
		Domain:   config.ControlAuthCookieDomain,
		Expires:  time.Unix(1, 0),
		MaxAge:   -1,
		Secure:   true,
		HTTPOnly: true,
		SameSite: config.ControlAuthCookieSameSite,
	})
}
