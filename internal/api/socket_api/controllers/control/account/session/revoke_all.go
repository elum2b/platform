package session

import (
	"strings"

	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type SessionRevokeAllResponse struct {
	Affected int64 `json:"affected"`
}

func RevokeAll(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		token := strings.TrimSpace(
			ctx.Peer.Session().GetAttribute(middleware.SessionTokenAttribute),
		)
		if token == "" {
			return serviceerrors.ErrUnauthorized
		}

		current, err := services.Control.Admin.ValidateSession(
			ctx,
			token,
			middleware.PeerIP(ctx.Peer),
		)
		if err != nil {
			return serviceerrors.ErrUnauthorized
		}

		affected, err := services.Control.Admin.RevokeAllSessions(
			ctx,
			ctx.Peer.Identity().UserID,
			current.ID,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			SessionRevokeAllResponse{Affected: affected},
		)
	})
}
