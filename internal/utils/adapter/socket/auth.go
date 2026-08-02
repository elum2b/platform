package socket

import (
	"context"
	"net"
	"strings"

	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
)

const sessionTokenAttribute = "control.session_token"

type sessionTokenContextKey struct{}

// Authenticated revalidates the server-side session before every event.
func Authenticated(next etp.Handler) etp.Handler {
	return func(ctx *etp.Context) error {
		if ctx == nil || ctx.Peer == nil {
			return serviceerrors.ErrUnauthorized
		}

		token := SessionToken(ctx.Peer)
		if token == "" {
			return serviceerrors.ErrUnauthorized
		}

		session, err := services.Control.Admin.ValidateSession(
			ctx,
			token,
			PeerIP(ctx.Peer),
		)
		if err != nil ||
			session.AccountID != strings.TrimSpace(ctx.Peer.Identity().UserID) {
			return serviceerrors.ErrUnauthorized
		}

		return next(ctx)
	}
}

func authHandler(
	ctx context.Context,
	peer *etp.Peer,
	_ etp.AuthRequest,
) (etp.AuthResult, error) {
	token := sessionToken(ctx)
	if token == "" {
		return etp.AuthResult{OK: false, Reason: "unauthorized"}, nil
	}

	session, err := services.Control.Admin.ValidateSession(
		ctx,
		token,
		PeerIP(peer),
	)
	if err != nil {
		//nolint:nilerr // ETP sends authentication denial in AuthResult.
		return etp.AuthResult{OK: false, Reason: "unauthorized"}, nil
	}

	return etp.AuthResult{
		OK:     true,
		UserID: session.AccountID,
		Attributes: []etp.AuthAttribute{{
			Key:   sessionTokenAttribute,
			Value: token,
		}},
	}, nil
}

// SessionToken returns the server-side token attached to an ETP session.
func SessionToken(peer *etp.Peer) string {
	if peer == nil {
		return ""
	}

	return strings.TrimSpace(
		peer.Session().GetAttribute(sessionTokenAttribute),
	)
}

// PeerIP normalizes both stream addresses and adapters that provide a bare IP.
func PeerIP(peer *etp.Peer) string {
	if peer == nil {
		return ""
	}

	remote := strings.TrimSpace(peer.RemoteAddr())

	host, _, err := net.SplitHostPort(remote)
	if err == nil {
		return host
	}

	return remote
}

func withSessionToken(ctx context.Context, token string) context.Context {
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
