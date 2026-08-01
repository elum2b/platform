package internalapi

import (
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestInitRegistersControlAuthenticationRoutes(t *testing.T) {
	app := fiber.New()
	Init(app)

	expected := map[string]string{
		"/internal/control/auth.check":         fiber.MethodGet,
		"/internal/control/auth.vkid":          fiber.MethodPost,
		"/internal/control/auth.telegram":      fiber.MethodPost,
		"/internal/control/auth.github":        fiber.MethodPost,
		"/internal/control/auth.gitlab":        fiber.MethodPost,
		"/internal/control/auth.google":        fiber.MethodPost,
		"/internal/control/auth.yandex":        fiber.MethodPost,
		"/internal/control/auth.ton.challenge": fiber.MethodGet,
		"/internal/control/auth.ton":           fiber.MethodPost,
		"/internal/control/auth.twoFactor":     fiber.MethodPost,
	}

	for _, route := range app.GetRoutes() {
		if method, ok := expected[route.Path]; ok && route.Method == method {
			delete(expected, route.Path)
		}
	}

	for path, method := range expected {
		t.Errorf("%s %s route is not registered", method, path)
	}
}
