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

type CreateRequest struct {
	WorkspaceID string     `json:"workspace_id" validate:"required,uuid"`
	RoleIDs     []string   `json:"role_ids"     validate:"dive,uuid"`
	ExpiresAt   *time.Time `json:"expires_at"`
}
type CreateResponse struct {
	ID          string     `json:"id"`
	Token       string     `json:"token"`
	WorkspaceID string     `json:"workspace_id"`
	ExpiresAt   *time.Time `json:"expires_at"`
	CreatedAt   time.Time  `json:"created_at"`
	RoleIDs     []string   `json:"role_ids"`
}

func Create(event string, socket etp.Router) {
	socket.Use(
		event,
		middleware.WorkspaceAccess("control.workspace.invite.create"),
	)
	socket.On(event, func(ctx *etp.Context) error {
		data := new(CreateRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		item, token, err := services.Control.Admin.CreateWorkspaceInvite(
			ctx,
			admin.CreateInviteParams{
				ActorID:     ctx.Peer.Identity().UserID,
				WorkspaceID: data.WorkspaceID,
				RoleIDs:     data.RoleIDs,
				ExpiresAt:   data.ExpiresAt,
			},
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			CreateResponse{
				ID:          item.ID,
				Token:       token,
				WorkspaceID: item.WorkspaceID,
				ExpiresAt:   item.ExpiresAt,
				CreatedAt:   item.CreatedAt,
				RoleIDs:     item.RoleIDs,
			},
		)
	})
}
