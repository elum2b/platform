package workspace

import (
	"time"

	etp "github.com/elum-utils/go-etp"
	controlmodel "github.com/elum2b/services/control/model"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type WorkspaceGetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type WorkspaceGetResponse struct {
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

func Get(event string, socket etp.Router) {
	socket.Use(event, middleware.WorkspaceAccess("control.workspace.get"))

	socket.On(event, func(ctx *etp.Context) error {
		data := new(WorkspaceGetRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		workspace, err := services.Control.Admin.GetWorkspace(
			ctx,
			data.WorkspaceID,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(ctx, event, WorkspaceGetResponse{
			ID:             workspace.ID,
			Slug:           workspace.Slug,
			Title:          workspace.Title,
			Status:         workspace.Status,
			CreatedBy:      workspace.CreatedBy,
			OwnerAccountID: workspace.OwnerAccountID,
			EmployeeLimit:  workspace.EmployeeLimit,
			CreatedAt:      workspace.CreatedAt,
			UpdatedAt:      workspace.UpdatedAt,
		})
	})
}
