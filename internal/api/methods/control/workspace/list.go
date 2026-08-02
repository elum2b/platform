package workspace

import (
	"time"

	controlmodel "github.com/elum2b/services/control/model"
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	Limit    int32     `json:"limit,omitempty"     validate:"omitempty,min=1,max=100"`
	CursorAt time.Time `json:"cursor_at,omitempty"`
	CursorID string    `json:"cursor_id,omitempty"`
}

type WorkspaceResponse struct {
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

type ListResponse struct {
	Workspaces []WorkspaceResponse `json:"workspaces"`
}

var (
	listKey         = "control.workspace.list"
	listDescription = `
Lists workspaces available to the current account.`
)

// List returns workspaces available to the authenticated account.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		items, err := services.Control.Admin.ListWorkspaces(
			ctx.Context,
			ctx.AccountID,
			controladmin.Page{
				Limit:    data.Limit,
				CursorAt: data.CursorAt,
				CursorID: data.CursorID,
			},
		)
		if err != nil {
			return ListResponse{}, err
		}

		workspaces := make([]WorkspaceResponse, 0, len(items))
		for _, item := range items {
			workspaces = append(workspaces, WorkspaceResponse{
				ID: item.ID, Slug: item.Slug, Title: item.Title,
				Status: item.Status, CreatedBy: item.CreatedBy,
				OwnerAccountID: item.OwnerAccountID,
				EmployeeLimit:  item.EmployeeLimit,
				CreatedAt:      item.CreatedAt, UpdatedAt: item.UpdatedAt,
			})
		}

		return ListResponse{Workspaces: workspaces}, nil
	},
}
