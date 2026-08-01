package workspace

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type MemberRemoveRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	AccountID   string `json:"account_id"   validate:"required,uuid"`
}

type MemberRemoveResponse struct {
	Affected int64 `json:"affected"`
}

func MemberRemove(event string, socket etp.Router) {
	socket.Use(
		event,
		middleware.WorkspaceAccess("control.workspace.member.remove"),
	)

	socket.On(event, func(ctx *etp.Context) error {
		data := new(MemberRemoveRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		affected, err := services.Control.Admin.RemoveMember(
			ctx,
			ctx.Peer.Identity().UserID,
			data.WorkspaceID,
			data.AccountID,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			MemberRemoveResponse{Affected: affected},
		)
	})
}
