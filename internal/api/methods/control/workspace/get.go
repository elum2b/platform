package workspace

import (
	"time"

	controlmodel "github.com/elum2b/services/control/model"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type GetResponse struct {
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
	getKey         = "control.workspace.get"
	getDescription = `
Returns detailed information about a workspace by ID. Requires the
'control.workspace.get' permission in the target workspace.`
)

// Get exposes the workspace retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.get"),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		workspace, err := services.Control.Admin.GetWorkspace(
			ctx.Context,
			data.WorkspaceID,
		)
		if err != nil {
			return GetResponse{}, err
		}

		return GetResponse{
			ID:             workspace.ID,
			Slug:           workspace.Slug,
			Title:          workspace.Title,
			Status:         workspace.Status,
			CreatedBy:      workspace.CreatedBy,
			OwnerAccountID: workspace.OwnerAccountID,
			EmployeeLimit:  workspace.EmployeeLimit,
			CreatedAt:      workspace.CreatedAt,
			UpdatedAt:      workspace.UpdatedAt,
		}, nil
	},
}
