package cookies

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
)

func TestSetUsesRequiredSecurityAttributes(t *testing.T) {
	originalDomain := config.ControlAuthCookieDomain
	originalSameSite := config.ControlAuthCookieSameSite

	t.Cleanup(func() {
		config.ControlAuthCookieDomain = originalDomain
		config.ControlAuthCookieSameSite = originalSameSite
	})

	config.ControlAuthCookieDomain = "example.com"
	config.ControlAuthCookieSameSite = "Strict"

	app := fiber.New()
	app.Get("/", func(ctx fiber.Ctx) error {
		Set(ctx, "session", "secret", time.Now().Add(time.Hour))

		return ctx.SendStatus(fiber.StatusNoContent)
	})

	response, err := app.Test(
		httptest.NewRequestWithContext(
			context.Background(),
			"GET",
			"https://example.com/",
			nil,
		),
	)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	cookie := strings.ToLower(response.Header.Get("Set-Cookie"))
	for _, attribute := range []string{
		"session=secret",
		"path=/",
		"domain=example.com",
		"httponly",
		"secure",
		"samesite=strict",
	} {
		if !strings.Contains(cookie, attribute) {
			t.Errorf("Set-Cookie does not contain %q: %s", attribute, cookie)
		}
	}
}
