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
		"/http/control.auth.check":         fiber.MethodGet,
		"/http/control.auth.provider.list": fiber.MethodGet,
		"/http/control.auth.vkid":          fiber.MethodPost,
		"/http/control.auth.telegram":      fiber.MethodPost,
		"/http/control.auth.discord":       fiber.MethodPost,
		"/http/control.auth.github":        fiber.MethodPost,
		"/http/control.auth.gitlab":        fiber.MethodPost,
		"/http/control.auth.google":        fiber.MethodPost,
		"/http/control.auth.yandex":        fiber.MethodPost,
		"/http/control.auth.ton.challenge": fiber.MethodGet,
		"/http/control.auth.ton":           fiber.MethodPost,
		"/http/control.auth.twoFactor":     fiber.MethodPost,
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

func TestInitRegistersUserGETAndPOSTRoutes(t *testing.T) {
	app := fiber.New()
	initHTTP(app)

	expectedPaths := []string{
		"/http/calendar.user.get",
		"/http/calendar.user.list_active",
		"/http/calendar.user.next",
		"/http/calendar.user.progress.get",
		"/http/calendar.user.record",
		"/http/cpa.user.code.get",
		"/http/cpa.user.offer.list",
		"/http/cpa.user.status.get",
		"/http/payment.user.create_attempt",
		"/http/payment.user.create_order",
		"/http/payment.user.create_order_by_key",
		"/http/payment.user.get_product",
		"/http/payment.user.get_product_by_key",
		"/http/payment.user.get_usdt_price",
		"/http/payment.user.is_subscription_active",
		"/http/payment.user.list_assets",
		"/http/payment.user.list_products",
		"/http/payment.user.list_usdt_prices",
		"/http/promo.apply",
		"/http/reference.user.get",
		"/http/reference.user.list",
		"/http/reference.user.resolve",
		"/http/tasks.user.claim",
		"/http/tasks.user.list_active",
		"/http/tasks.user.partner.check",
		"/http/tasks.user.partner.list",
		"/http/tasks.user.partner.start",
		"/http/tasks.user.start",
	}

	expected := make(map[string]map[string]struct{}, len(expectedPaths))
	for _, path := range expectedPaths {
		expected[path] = map[string]struct{}{
			fiber.MethodGet:  {},
			fiber.MethodPost: {},
		}
	}

	for _, route := range app.GetRoutes() {
		if methods, ok := expected[route.Path]; ok {
			delete(methods, route.Method)
		}
	}

	for path, methods := range expected {
		for method := range methods {
			t.Errorf("%s %s route is not registered", method, path)
		}
	}
}

func TestAuthCheckAcceptsGETWithoutBody(t *testing.T) {
	app := fiber.New()
	initHTTP(app)

	response, err := app.Test(httptest.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"/http/control.auth.check",
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
