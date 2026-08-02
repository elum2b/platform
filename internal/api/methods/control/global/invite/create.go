package invite

import (
	"time"

	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateRequest struct {
	RoleIDs   []string   `json:"role_ids"             validate:"dive,uuid"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

type CreateResponse struct {
	ID        string     `json:"id"`
	Token     string     `json:"token"`
	ExpiresAt *time.Time `json:"expires_at"`
	CreatedAt time.Time  `json:"created_at"`
	RoleIDs   []string   `json:"role_ids"`
}

var (
	createKey         = "control.global.invite.create"
	createDescription = `
Creates an invitation to join the platform. Requires the
'control.global.invite.create' global permission.`
)

// Create creates a global platform invitation.
var Create = adapter.Method[CreateRequest, CreateResponse]{
	Key:         createKey,
	Description: createDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.invite.create"),
	},
	Handler: func(ctx *adapter.Context, data CreateRequest) (CreateResponse, error) {
		item, token, err := services.Control.Admin.CreateGlobalInvite(
			ctx.Context,
			controladmin.CreateInviteParams{
				ActorID: ctx.AccountID, RoleIDs: data.RoleIDs,
				ExpiresAt: data.ExpiresAt,
			},
		)
		if err != nil {
			return CreateResponse{}, err
		}

		return CreateResponse{
			ID: item.ID, Token: token, ExpiresAt: item.ExpiresAt,
			CreatedAt: item.CreatedAt, RoleIDs: item.RoleIDs,
		}, nil
	},
}
