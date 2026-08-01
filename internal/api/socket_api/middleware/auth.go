package middleware

import (
	"net"
	"strings"

	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
)

// SessionTokenAttribute is kept only in the server-side ETP session identity.
// It is used to revalidate a connection before every protected event.
const SessionTokenAttribute = "control.session_token"

// Authenticated rejects an event when its session was revoked, expired, became
// inactive, or no longer matches an IP-bound session. Register it on a route
// group before access-control middleware.
func Authenticated(next etp.Handler) etp.Handler {
	return func(ctx *etp.Context) error {
		if ctx == nil || ctx.Peer == nil {
			return serviceerrors.ErrUnauthorized
		}

		if services.Control == nil || !services.Control.IsReady() {
			return serviceerrors.ErrNotReady
		}

		token := strings.TrimSpace(
			ctx.Peer.Session().GetAttribute(SessionTokenAttribute),
		)
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
