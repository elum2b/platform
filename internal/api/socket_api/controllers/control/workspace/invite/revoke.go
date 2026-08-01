package invite

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type RevokeRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	InviteID    string `json:"invite_id"    validate:"required,uuid"`
}
type RevokeResponse struct {
	Affected int64 `json:"affected"`
}

func Revoke(event string, socket etp.Router) {
	socket.Use(
		event,
		middleware.WorkspaceAccess("control.workspace.invite.revoke"),
	)
	socket.On(event, func(ctx *etp.Context) error {
		data := new(RevokeRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		affected, err := services.Control.Admin.RevokeInvite(
			ctx,
			ctx.Peer.Identity().UserID,
			data.InviteID,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			RevokeResponse{Affected: affected},
		)
	})
}
