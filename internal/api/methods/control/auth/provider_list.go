package auth

import (
	"github.com/gofiber/fiber/v3"

	adapter "github.com/elum2b/platform/internal/utils/adapter"
	authutils "github.com/elum2b/platform/internal/utils/auth"
)

var (
	providerListKey         = "control.auth.provider.list"
	providerListDescription = `
Lists authentication providers configured for frontend use.`
)

// ProviderList returns the authentication providers available to the frontend.
var ProviderList = adapter.Method[struct{}, []authutils.Provider]{
	Key:         providerListKey,
	Description: providerListDescription,
	Transports:  adapter.HTTP,
	Method:      fiber.MethodGet,
	Handler: func(
		_ *adapter.Context,
		_ struct{},
	) ([]authutils.Provider, error) {
		return authutils.Providers(), nil
	},
}
