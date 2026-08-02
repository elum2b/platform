package mcp

import (
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type RevokeRequest struct {
	TokenID string `json:"token_id" validate:"required,uuid"`
}

type RevokeResponse struct {
	Affected int64 `json:"affected"`
}

var (
	revokeKey         = "control.account.mcp.revoke"
	revokeDescription = `
Revokes an MCP token owned by the authenticated account.`
)

// Revoke revokes an MCP token owned by the authenticated account.
var Revoke = adapter.Method[RevokeRequest, RevokeResponse]{
	Key:         revokeKey,
	Description: revokeDescription,
	Transports:  adapter.WS,
	Handler: func(
		ctx *adapter.Context,
		data RevokeRequest,
	) (RevokeResponse, error) {
		affected, err := services.Control.Admin.RevokeMCPToken(
			ctx.Context,
			controladmin.RevokeMCPTokenParams{
				AccountID: ctx.AccountID,
				TokenID:   data.TokenID,
			},
		)
		if err != nil {
			return RevokeResponse{}, err
		}

		return RevokeResponse{Affected: affected}, nil
	},
}
