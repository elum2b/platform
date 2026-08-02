package workspace

import (
	"time"

	controlmodel "github.com/elum2b/services/control/model"
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateRequest struct {
	ID    string `json:"id,omitempty" validate:"omitempty,uuid"`
	Slug  string `json:"slug"         validate:"required,max=255"`
	Title string `json:"title"        validate:"required,max=255"`
}

type CreateResponse struct {
	ID             string                       `json:"id"`
	Slug           string                       `json:"slug"`
	Title          string                       `json:"title"`
	Status         controlmodel.WorkspaceStatus `json:"status"`
	CreatedBy      string                       `json:"created_by"`
	OwnerAccountID string                       `json:"owner_account_id"`
	EmployeeLimit  int32                        `json:"employee_limit"`
	CreatedAt      time.Time                    `json:"created_at"`
	UpdatedAt      time.Time                    `json:"updated_at"`
}

var (
	createKey         = "control.global.workspace.create"
	createDescription = `
Creates a workspace owned by the current account. Requires the
'control.global.workspace.create' global permission.`
)

// Create creates a workspace owned by the authenticated account.
var Create = adapter.Method[CreateRequest, CreateResponse]{
	Key:         createKey,
	Description: createDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.workspace.create"),
	},
	Handler: func(ctx *adapter.Context, data CreateRequest) (CreateResponse, error) {
		workspace, err := services.Control.Admin.CreateWorkspace(
			ctx.Context,
			controladmin.CreateWorkspaceParams{
				ActorID: ctx.AccountID,
				ID:      data.ID, Slug: data.Slug, Title: data.Title,
			},
		)
		if err != nil {
			return CreateResponse{}, err
		}

		return CreateResponse{
			ID: workspace.ID, Slug: workspace.Slug, Title: workspace.Title,
			Status: workspace.Status, CreatedBy: workspace.CreatedBy,
			OwnerAccountID: workspace.OwnerAccountID,
			EmployeeLimit:  workspace.EmployeeLimit,
			CreatedAt:      workspace.CreatedAt, UpdatedAt: workspace.UpdatedAt,
		}, nil
	},
}
