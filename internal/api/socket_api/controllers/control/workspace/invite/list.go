package invite

import (
	"time"

	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type ListRequest struct {
	WorkspaceID string    `json:"workspace_id" validate:"required,uuid"`
	Limit       int32     `json:"limit"        validate:"omitempty,min=1,max=100"`
	CursorAt    time.Time `json:"cursor_at"`
	CursorID    string    `json:"cursor_id"`
}
type ListResponse struct {
	ID          string     `json:"id"`
	WorkspaceID string     `json:"workspace_id"`
	CreatedBy   string     `json:"created_by"`
	ExpiresAt   *time.Time `json:"expires_at"`
	AcceptedBy  string     `json:"accepted_by"`
	AcceptedAt  *time.Time `json:"accepted_at"`
	RevokedAt   *time.Time `json:"revoked_at"`
	CreatedAt   time.Time  `json:"created_at"`
	RoleIDs     []string   `json:"role_ids"`
}

func List(event string, socket etp.Router) {
	socket.Use(
		event,
		middleware.WorkspaceAccess("control.workspace.invite.list"),
	)
	socket.On(event, func(ctx *etp.Context) error {
		data := new(ListRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		items, err := services.Control.Admin.ListWorkspaceInvites(
			ctx,
			data.WorkspaceID,
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
					ID:          item.ID,
					WorkspaceID: item.WorkspaceID,
					CreatedBy:   item.CreatedBy,
					ExpiresAt:   item.ExpiresAt,
					AcceptedBy:  item.AcceptedBy,
					AcceptedAt:  item.AcceptedAt,
					RevokedAt:   item.RevokedAt,
					CreatedAt:   item.CreatedAt,
					RoleIDs:     item.RoleIDs,
				},
			)
		}

		return socketutils.Respond(ctx, event, response)
	})
}
