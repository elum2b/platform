package mcp

import (
	"time"

	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListResponse struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	ExpiresAt  *time.Time `json:"expires_at"`
	RevokedAt  *time.Time `json:"revoked_at"`
	LastUsedAt time.Time  `json:"last_used_at"`
	CreatedAt  time.Time  `json:"created_at"`
}

var (
	listKey         = "control.account.mcp.list"
	listDescription = `
Lists MCP tokens issued by the authenticated account without their secrets.`
)

// List returns MCP tokens issued by the authenticated account.
var List = adapter.Method[struct{}, []ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, _ struct{}) ([]ListResponse, error) {
		tokens, err := services.Control.Admin.ListMCPTokens(
			ctx.Context,
			controladmin.ListMCPTokensParams{AccountID: ctx.AccountID},
		)
		if err != nil {
			return nil, err
		}

		response := make([]ListResponse, 0, len(tokens))
		for _, token := range tokens {
			response = append(response, ListResponse{
				ID:         token.ID,
				Name:       token.Name,
				ExpiresAt:  token.ExpiresAt,
				RevokedAt:  token.RevokedAt,
				LastUsedAt: token.LastUsedAt,
				CreatedAt:  token.CreatedAt,
			})
		}

		return response, nil
	},
}
