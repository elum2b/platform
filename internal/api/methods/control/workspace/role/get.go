package role

import (
	"time"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	RoleID      string `json:"role_id"      validate:"required,uuid"`
}

type GetResponse struct {
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

var (
	getKey         = "control.workspace.role.get"
	getDescription = `
Gets a workspace role by ID. Requires the
'control.workspace.role.get' permission in the target workspace.`
)

// Get returns a workspace role.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.role.get"),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		item, err := services.Control.Admin.GetWorkspaceRole(
			ctx.Context,
			data.WorkspaceID,
			data.RoleID,
		)
		if err != nil {
			return GetResponse{}, err
		}

		return GetResponse{
			ID: item.ID, WorkspaceID: item.WorkspaceID, Code: item.Code,
			Title: item.Title, Description: item.Description,
			Position: item.Position, MemberCount: item.MemberCount,
			CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt,
		}, nil
	},
}
