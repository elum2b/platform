package member

import (
	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type AssignRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	AccountID   string `json:"account_id"   validate:"required,uuid"`
	RoleID      string `json:"role_id"      validate:"required,uuid"`
}

func Assign(event string, socket etp.Router) {
	socket.Use(
		event,
		middleware.WorkspaceAccess("control.workspace.role.member.assign"),
	)
	socket.On(event, func(ctx *etp.Context) error {
		data := new(AssignRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		if err := services.Control.Admin.AssignWorkspaceRole(
			ctx,
			admin.SetRoleMemberParams{
				ActorID:     ctx.Peer.Identity().UserID,
				WorkspaceID: data.WorkspaceID,
				AccountID:   data.AccountID,
				RoleID:      data.RoleID,
			},
		); err != nil {
			return err
		}

		return socketutils.Respond(ctx, event, struct{}{})
	})
}
