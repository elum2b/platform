package invite

import (
	"time"

	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string    `json:"workspace_id"        validate:"required,uuid"`
	Limit       int32     `json:"limit,omitempty"     validate:"omitempty,min=1,max=100"`
	CursorAt    time.Time `json:"cursor_at,omitempty"`
	CursorID    string    `json:"cursor_id,omitempty"`
}

type Item struct {
	ID          string     `json:"id"`
	WorkspaceID string     `json:"workspace_id"`
	CreatedBy   string     `json:"created_by"`
	ExpiresAt   *time.Time `json:"expires_at"`
	AcceptedBy  string     `json:"accepted_by"`
	AcceptedAt  *time.Time `json:"accepted_at"`
	RevokedAt   *time.Time `json:"revoked_at"`
	CreatedAt   time.Time  `json:"created_at"`
	RoleIDs     []string   `json:"role_ids"`
}

type ListResponse struct {
	Invites []Item `json:"invites"`
}

var (
	listKey         = "control.workspace.invite.list"
	listDescription = `
Lists invitations created for a workspace. Requires the
'control.workspace.invite.list' permission in the target workspace.`
)

// List returns workspace invitations.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.invite.list"),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		items, err := services.Control.Admin.ListWorkspaceInvites(
			ctx.Context,
			data.WorkspaceID,
			controladmin.Page{
				Limit:    data.Limit,
				CursorAt: data.CursorAt,
				CursorID: data.CursorID,
			},
		)
		if err != nil {
			return ListResponse{}, err
		}

		invites := make([]Item, 0, len(items))
		for _, item := range items {
			invites = append(invites, Item{
				ID: item.ID, WorkspaceID: item.WorkspaceID,
				CreatedBy: item.CreatedBy, ExpiresAt: item.ExpiresAt,
				AcceptedBy: item.AcceptedBy, AcceptedAt: item.AcceptedAt,
				RevokedAt: item.RevokedAt, CreatedAt: item.CreatedAt,
				RoleIDs: item.RoleIDs,
			})
		}

		return ListResponse{Invites: invites}, nil
	},
}
