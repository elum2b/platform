package workspace

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type OwnerTransferRequest struct {
	WorkspaceID     string `json:"workspace_id"      validate:"required,uuid"`
	TargetAccountID string `json:"target_account_id" validate:"required,uuid"`
}

type OwnerTransferResponse struct {
	Transferred bool `json:"transferred"`
}

func OwnerTransfer(event string, socket etp.Router) {
	socket.Use(
		event,
		middleware.WorkspaceAccess("control.workspace.owner.transfer"),
	)

	socket.On(event, func(ctx *etp.Context) error {
		data := new(OwnerTransferRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		err := services.Control.Admin.TransferWorkspaceOwnership(
			ctx,
			ctx.Peer.Identity().UserID,
			data.WorkspaceID,
			data.TargetAccountID,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			OwnerTransferResponse{Transferred: true},
		)
	})
}
