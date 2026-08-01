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

type EmployeeLimitRequestRequest struct {
	WorkspaceID    string `json:"workspace_id"    validate:"required,uuid"`
	RequestedLimit int32  `json:"requested_limit" validate:"required,min=1"`
	Reason         string `json:"reason"          validate:"max=1000"`
}
type EmployeeLimitRequestResponse struct {
	ID             string                          `json:"id"`
	Kind           admin.LimitKind                 `json:"kind"`
	WorkspaceID    string                          `json:"workspace_id"`
	CurrentLimit   int32                           `json:"current_limit"`
	RequestedLimit int32                           `json:"requested_limit"`
	Reason         string                          `json:"reason"`
	Status         controlmodel.LimitRequestStatus `json:"status"`
	CreatedAt      time.Time                       `json:"created_at"`
}

func EmployeeLimitRequest(event string, socket etp.Router) {
	socket.Use(
		event,
		middleware.WorkspaceAccess("control.workspace.employee_limit.request"),
	)
	socket.On(event, func(ctx *etp.Context) error {
		data := new(EmployeeLimitRequestRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		item, err := services.Control.Admin.RequestEmployeeLimit(
			ctx,
			ctx.Peer.Identity().UserID,
			data.WorkspaceID,
			data.RequestedLimit,
			data.Reason,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			EmployeeLimitRequestResponse{
				ID:             item.ID,
				Kind:           item.Kind,
				WorkspaceID:    item.WorkspaceID,
				CurrentLimit:   item.CurrentLimit,
				RequestedLimit: item.RequestedLimit,
				Reason:         item.Reason,
				Status:         item.Status,
				CreatedAt:      item.CreatedAt,
			},
		)
	})
}
