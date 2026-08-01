package system

import (
	"context"
	"strings"

	etp "github.com/elum-utils/go-etp"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
)

type sessionTokenContextKey struct{}

// AuthHandler validates the session token supplied by the HTTP-only cookie.
// The accepted token stays server-side and is revalidated by Authenticated for
// every routed event.
func AuthHandler(
	ctx context.Context,
	peer *etp.Peer,
	_ etp.AuthRequest,
) (etp.AuthResult, error) {
	if services.Control == nil || !services.Control.IsReady() {
		return etp.AuthResult{OK: false, Reason: "service unavailable"}, nil
	}

	token := sessionToken(ctx)
	if token == "" {
		return etp.AuthResult{OK: false, Reason: "unauthorized"}, nil
	}

	session, err := services.Control.Admin.ValidateSession(
		ctx,
		token,
		middleware.PeerIP(peer),
	)
	if err != nil {
		//nolint:nilerr // ETP sends authentication denial in AuthResult, not as a transport error.
		return etp.AuthResult{OK: false, Reason: "unauthorized"}, nil
	}

	return etp.AuthResult{
		OK:     true,
		UserID: session.AccountID,
		Attributes: []etp.AuthAttribute{{
			Key:   middleware.SessionTokenAttribute,
			Value: token,
		}},
	}, nil
}

func WithSessionToken(ctx context.Context, token string) context.Context {
	return context.WithValue(
		ctx,
		sessionTokenContextKey{},
		strings.TrimSpace(token),
	)
}

func sessionToken(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	token, _ := ctx.Value(sessionTokenContextKey{}).(string)

	return strings.TrimSpace(token)
}
