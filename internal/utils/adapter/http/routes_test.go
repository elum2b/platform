package http_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/api/methods"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	httputils "github.com/elum2b/platform/internal/utils/adapter/http"
)

func initHTTP(app fiber.Router) {
	httputils.Init(app, func(router fiber.Router) {
		methods.Register(adapter.Registry{HTTP: router})
	})
}

func TestInitRegistersControlAuthenticationRoutes(t *testing.T) {
	app := fiber.New()
	initHTTP(app)

	expected := map[string]string{
		"/internal/control/control.auth.check":         fiber.MethodGet,
		"/internal/control/control.auth.vkid":          fiber.MethodPost,
		"/internal/control/control.auth.telegram":      fiber.MethodPost,
		"/internal/control/control.auth.github":        fiber.MethodPost,
		"/internal/control/control.auth.gitlab":        fiber.MethodPost,
		"/internal/control/control.auth.google":        fiber.MethodPost,
		"/internal/control/control.auth.yandex":        fiber.MethodPost,
		"/internal/control/control.auth.ton.challenge": fiber.MethodGet,
		"/internal/control/control.auth.ton":           fiber.MethodPost,
		"/internal/control/control.auth.twoFactor":     fiber.MethodPost,
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

func TestAuthCheckAcceptsGETWithoutBody(t *testing.T) {
	app := fiber.New()
	initHTTP(app)

	response, err := app.Test(httptest.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"/internal/control/control.auth.check",
		http.NoBody,
	))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf(
			"GET auth check status = %d, want %d",
			response.StatusCode,
			http.StatusOK,
		)
	}
}
