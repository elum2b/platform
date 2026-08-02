package role

import (
	"time"

	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateRequest struct {
	WorkspaceID string `json:"workspace_id"          validate:"required,uuid"`
	ID          string `json:"id,omitempty"          validate:"omitempty,uuid"`
	Code        string `json:"code"                  validate:"required,max=255"`
	Title       string `json:"title"                 validate:"required,max=255"`
	Description string `json:"description,omitempty" validate:"max=1000"`
	Position    int32  `json:"position,omitempty"`
}

type CreateResponse struct {
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
	createKey         = "control.workspace.role.create"
	createDescription = `
Creates a role in a workspace. Requires the
'control.workspace.role.create' permission in the target workspace.`
)

// Create creates a workspace role.
var Create = adapter.Method[CreateRequest, CreateResponse]{
	Key:         createKey,
	Description: createDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.role.create"),
	},
	Handler: func(ctx *adapter.Context, data CreateRequest) (CreateResponse, error) {
		item, err := services.Control.Admin.CreateWorkspaceRole(
			ctx.Context,
			controladmin.CreateRoleParams{
				ActorID: ctx.AccountID, WorkspaceID: data.WorkspaceID,
				ID: data.ID, Code: data.Code, Title: data.Title,
				Description: data.Description, Position: data.Position,
			},
		)
		if err != nil {
			return CreateResponse{}, err
		}

		return CreateResponse{
			ID: item.ID, WorkspaceID: item.WorkspaceID, Code: item.Code,
			Title: item.Title, Description: item.Description,
			Position: item.Position, MemberCount: item.MemberCount,
			CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt,
		}, nil
	},
}
