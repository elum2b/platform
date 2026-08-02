package session

import (
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	socketutils "github.com/elum2b/platform/internal/utils/adapter/socket"
)

type RevokeAllResponse struct {
	Affected int64 `json:"affected"`
}

var (
	revokeAllKey         = "control.account.session.revoke_all"
	revokeAllDescription = `
Revokes all account sessions except the current WebSocket session.`
)

// RevokeAll revokes every session except the current WebSocket session.
var RevokeAll = adapter.Method[struct{}, RevokeAllResponse]{
	Key:         revokeAllKey,
	Description: revokeAllDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, _ struct{}) (RevokeAllResponse, error) {
		token := socketutils.SessionToken(ctx.Socket.Peer)
		if token == "" {
			return RevokeAllResponse{}, serviceerrors.ErrUnauthorized
		}

		current, err := services.Control.Admin.ValidateSession(
			ctx.Context,
			token,
			socketutils.PeerIP(ctx.Socket.Peer),
		)
		if err != nil {
			return RevokeAllResponse{}, serviceerrors.ErrUnauthorized
		}

		affected, err := services.Control.Admin.RevokeAllSessions(
			ctx.Context,
			ctx.AccountID,
			current.ID,
		)
		if err != nil {
			return RevokeAllResponse{}, err
		}

		return RevokeAllResponse{Affected: affected}, nil
	},
}
