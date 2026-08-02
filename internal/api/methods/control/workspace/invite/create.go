package invite

import (
	"time"

	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateRequest struct {
	WorkspaceID string     `json:"workspace_id"         validate:"required,uuid"`
	RoleIDs     []string   `json:"role_ids"             validate:"dive,uuid"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
}

type CreateResponse struct {
	ID          string     `json:"id"`
	Token       string     `json:"token"`
	WorkspaceID string     `json:"workspace_id"`
	ExpiresAt   *time.Time `json:"expires_at"`
	CreatedAt   time.Time  `json:"created_at"`
	RoleIDs     []string   `json:"role_ids"`
}

var (
	createKey         = "control.workspace.invite.create"
	createDescription = `
Creates an invitation for a workspace. Requires the
'control.workspace.invite.create' permission in the target workspace.`
)

// Create creates a workspace invitation.
var Create = adapter.Method[CreateRequest, CreateResponse]{
	Key:         createKey,
	Description: createDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.invite.create"),
	},
	Handler: func(ctx *adapter.Context, data CreateRequest) (CreateResponse, error) {
		item, token, err := services.Control.Admin.CreateWorkspaceInvite(
			ctx.Context,
			controladmin.CreateInviteParams{
				ActorID:     ctx.AccountID,
				WorkspaceID: data.WorkspaceID,
				RoleIDs:     data.RoleIDs,
				ExpiresAt:   data.ExpiresAt,
			},
		)
		if err != nil {
			return CreateResponse{}, err
		}

		return CreateResponse{
			ID: item.ID, Token: token, WorkspaceID: item.WorkspaceID,
			ExpiresAt: item.ExpiresAt, CreatedAt: item.CreatedAt,
			RoleIDs: item.RoleIDs,
		}, nil
	},
}
