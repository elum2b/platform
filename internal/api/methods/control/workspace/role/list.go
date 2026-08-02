package role

import (
	"time"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type Item struct {
	ID          string    `json:"id"`
	WorkspaceID string    `json:"workspace_id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Position    int32     `json:"position"`
	MemberCount int64     `json:"member_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListResponse struct {
	Roles []Item `json:"roles"`
}

var (
	listKey         = "control.workspace.role.list"
	listDescription = `
Lists roles configured for a workspace. Requires the
'control.workspace.role.list' permission in the target workspace.`
)

// List returns workspace roles.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.role.list"),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		items, err := services.Control.Admin.ListWorkspaceRoles(
			ctx.Context,
			data.WorkspaceID,
		)
		if err != nil {
			return ListResponse{}, err
		}

		roles := make([]Item, 0, len(items))
		for _, item := range items {
			roles = append(roles, Item{
				ID: item.ID, WorkspaceID: item.WorkspaceID, Code: item.Code,
				Title: item.Title, Description: item.Description,
				Position: item.Position, MemberCount: item.MemberCount,
				CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt,
			})
		}

		return ListResponse{Roles: roles}, nil
	},
}
