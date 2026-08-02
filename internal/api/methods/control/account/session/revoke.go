package session

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type RevokeRequest struct {
	SessionID string `json:"session_id" validate:"required,uuid"`
}

type RevokeResponse struct {
	Affected int64 `json:"affected"`
}

var (
	revokeKey         = "control.account.session.revoke"
	revokeDescription = `
Revokes one session owned by the authenticated account.`
)

// Revoke revokes one session owned by the authenticated account.
var Revoke = adapter.Method[RevokeRequest, RevokeResponse]{
	Key:         revokeKey,
	Description: revokeDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, data RevokeRequest) (RevokeResponse, error) {
		affected, err := services.Control.Admin.RevokeSession(
			ctx.Context,
			ctx.AccountID,
			data.SessionID,
		)
		if err != nil {
			return RevokeResponse{}, err
		}

		return RevokeResponse{Affected: affected}, nil
	},
}
