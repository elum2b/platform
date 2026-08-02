package invite

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type RevokeRequest struct {
	InviteID string `json:"invite_id" validate:"required,uuid"`
}

type RevokeResponse struct {
	Affected int64 `json:"affected"`
}

var (
	revokeKey         = "control.global.invite.revoke"
	revokeDescription = `
Revokes an active platform invitation. Requires the
'control.global.invite.revoke' global permission.`
)

// Revoke revokes a global platform invitation.
var Revoke = adapter.Method[RevokeRequest, RevokeResponse]{
	Key:         revokeKey,
	Description: revokeDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.invite.revoke"),
	},
	Handler: func(ctx *adapter.Context, data RevokeRequest) (RevokeResponse, error) {
		affected, err := services.Control.Admin.RevokeInvite(
			ctx.Context,
			ctx.AccountID,
			data.InviteID,
		)
		if err != nil {
			return RevokeResponse{}, err
		}

		return RevokeResponse{Affected: affected}, nil
	},
}
