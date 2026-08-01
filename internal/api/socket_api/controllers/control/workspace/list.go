package workspace

import (
	"time"

	etp "github.com/elum-utils/go-etp"
	controlmodel "github.com/elum2b/services/control/model"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type ListRequest struct {
	Limit    int32     `json:"limit"     validate:"omitempty,min=1,max=100"`
	CursorAt time.Time `json:"cursor_at"`
	CursorID string    `json:"cursor_id"`
}

type ListResponse struct {
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

func List(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		data := new(ListRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		workspaces, err := services.Control.Admin.ListWorkspaces(
			ctx,
			ctx.Peer.Identity().UserID,
			admin.Page{
				Limit:    data.Limit,
				CursorAt: data.CursorAt,
				CursorID: data.CursorID,
			},
		)
		if err != nil {
			return err
		}

		response := make([]ListResponse, 0, len(workspaces))
		for _, workspace := range workspaces {
			response = append(response, ListResponse{
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
		}

		return socketutils.Respond(ctx, event, response)
	})
}
