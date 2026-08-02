package invite

import (
	"time"

	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	Limit    int32     `json:"limit,omitempty"     validate:"omitempty,min=1,max=100"`
	CursorAt time.Time `json:"cursor_at,omitempty"`
	CursorID string    `json:"cursor_id,omitempty"`
}

type Item struct {
	ID         string     `json:"id"`
	CreatedBy  string     `json:"created_by"`
	ExpiresAt  *time.Time `json:"expires_at"`
	AcceptedBy string     `json:"accepted_by"`
	AcceptedAt *time.Time `json:"accepted_at"`
	RevokedAt  *time.Time `json:"revoked_at"`
	CreatedAt  time.Time  `json:"created_at"`
	RoleIDs    []string   `json:"role_ids"`
}

type ListResponse struct {
	Invites []Item `json:"invites"`
}

var (
	listKey         = "control.global.invite.list"
	listDescription = `
Lists invitations created for the platform. Requires the
'control.global.invite.list' global permission.`
)

// List returns global platform invitations.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.invite.list"),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		items, err := services.Control.Admin.ListGlobalInvites(
			ctx.Context,
			controladmin.Page{
				Limit: data.Limit, CursorAt: data.CursorAt,
				CursorID: data.CursorID,
			},
		)
		if err != nil {
			return ListResponse{}, err
		}

		invites := make([]Item, 0, len(items))
		for _, item := range items {
			invites = append(invites, Item{
				ID: item.ID, CreatedBy: item.CreatedBy,
				ExpiresAt: item.ExpiresAt, AcceptedBy: item.AcceptedBy,
				AcceptedAt: item.AcceptedAt, RevokedAt: item.RevokedAt,
				CreatedAt: item.CreatedAt, RoleIDs: item.RoleIDs,
			})
		}

		return ListResponse{Invites: invites}, nil
	},
}
