package member

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

type ListRequest struct {
	Limit    int32     `json:"limit"     validate:"omitempty,min=1,max=100"`
	CursorAt time.Time `json:"cursor_at"`
	CursorID string    `json:"cursor_id"`
}
type ListResponse struct {
	AccountID           string                        `json:"account_id"`
	DisplayName         string                        `json:"display_name"`
	Status              controlmodel.MembershipStatus `json:"status"`
	WorkspaceLimit      int32                         `json:"workspace_limit"`
	OwnedWorkspaceCount int64                         `json:"owned_workspace_count"`
	InvitedBy           string                        `json:"invited_by"`
	JoinedAt            time.Time                     `json:"joined_at"`
	UpdatedAt           time.Time                     `json:"updated_at"`
}

func List(event string, socket etp.Router) {
	socket.Use(event, middleware.GlobalAccess("control.global.member.list"))
	socket.On(event, func(ctx *etp.Context) error {
		data := new(ListRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		items, err := services.Control.Admin.ListPlatformMembers(
			ctx,
			admin.Page{
				Limit:    data.Limit,
				CursorAt: data.CursorAt,
				CursorID: data.CursorID,
			},
		)
		if err != nil {
			return err
		}

		response := make([]ListResponse, 0, len(items))
		for _, item := range items {
			response = append(
				response,
				ListResponse{
					AccountID:           item.AccountID,
					DisplayName:         item.DisplayName,
					Status:              item.Status,
					WorkspaceLimit:      item.WorkspaceLimit,
					OwnedWorkspaceCount: item.OwnedWorkspaceCount,
					InvitedBy:           item.InvitedBy,
					JoinedAt:            item.JoinedAt,
					UpdatedAt:           item.UpdatedAt,
				},
			)
		}

		return socketutils.Respond(ctx, event, response)
	})
}
