package mcp

import (
	"context"
	"strings"

	controlinternal "github.com/elum2b/services/control/service/internalapi"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
)

type tokenContextKey struct{}

// TokenFromContext returns the MCP bearer token extracted from the request.
func TokenFromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(tokenContextKey{}).(string)
	if !ok || strings.TrimSpace(token) == "" {
		return "", false
	}

	return token, true
}

func withToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenContextKey{}, strings.TrimSpace(token))
}

// Authenticated validates the MCP token before a tool handler runs.
func Authenticated(handler Handler) {
	handler.Use(func(ctx context.Context) (context.Context, error) {
		token, ok := TokenFromContext(ctx)
		if !ok {
			return ctx, serviceerrors.ErrUnauthorized
		}

		principal, err := services.Control.Internal.ValidateMCPToken(
			ctx,
			controlinternal.ValidateMCPTokenRequest{Token: token},
		)
		if err != nil {
			return ctx, serviceerrors.ErrUnauthorized
		}

		return WithPrincipal(ctx, principal), nil
	})
}
