package workspace

import (
	"time"

	etp "github.com/elum-utils/go-etp"
	controlmodel "github.com/elum2b/services/control/model"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type CreateRequest struct {
	ID    string `json:"id"    validate:"omitempty,uuid"`
	Slug  string `json:"slug"  validate:"required,max=255"`
	Title string `json:"title" validate:"required,max=255"`
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

func Create(event string, socket etp.Router) {
	socket.Use(
		event,
		middleware.GlobalAccess("control.global.workspace.create"),
	)

	socket.On(event, func(ctx *etp.Context) error {
		data := new(CreateRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		workspace, err := services.Control.Admin.CreateWorkspace(
			ctx,
			admin.CreateWorkspaceParams{
				ActorID: ctx.Peer.Identity().UserID,
				ID:      data.ID,
				Slug:    data.Slug,
				Title:   data.Title,
			},
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(ctx, event, CreateResponse{
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
