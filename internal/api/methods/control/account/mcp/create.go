package mcp

import (
	"time"

	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateRequest struct {
	Name     string `json:"name"     validate:"required,max=128"`
	Duration int64  `json:"duration" validate:"min=0,max=9223372036854" jsonschema:"Token lifetime in milliseconds. Zero means no expiration."`
}

type CreateTokenResponse struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	ExpiresAt  *time.Time `json:"expires_at"`
	RevokedAt  *time.Time `json:"revoked_at"`
	LastUsedAt time.Time  `json:"last_used_at"`
	CreatedAt  time.Time  `json:"created_at"`
}

type CreateResponse struct {
	Token    CreateTokenResponse `json:"token"`
	RawToken string              `json:"raw_token"`
}

var (
	createKey         = "control.account.mcp.create"
	createDescription = `
Creates an MCP token for the authenticated account and returns its secret once.`
)

// Create issues an MCP token for the authenticated account.
var Create = adapter.Method[CreateRequest, CreateResponse]{
	Key:         createKey,
	Description: createDescription,
	Transports:  adapter.WS,
	Handler: func(
		ctx *adapter.Context,
		data CreateRequest,
	) (CreateResponse, error) {
		result, err := services.Control.Admin.CreateMCPToken(
			ctx.Context,
			controladmin.CreateMCPTokenParams{
				AccountID: ctx.AccountID,
				Name:      data.Name,
				Duration:  time.Duration(data.Duration) * time.Millisecond,
			},
		)
		if err != nil {
			return CreateResponse{}, err
		}

		return CreateResponse{
			Token: CreateTokenResponse{
				ID:         result.Token.ID,
				Name:       result.Token.Name,
				ExpiresAt:  result.Token.ExpiresAt,
				RevokedAt:  result.Token.RevokedAt,
				LastUsedAt: result.Token.LastUsedAt,
				CreatedAt:  result.Token.CreatedAt,
			},
			RawToken: result.RawToken,
		}, nil
	},
}
