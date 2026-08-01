package global

import (
	"github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type OwnerTransferRequest struct {
	AccountID string `json:"account_id" validate:"required,uuid"`
}

func OwnerTransfer(event string, socket etp.Router) {
	socket.Use(event, middleware.GlobalAccess("control.global.owner.transfer"))
	socket.On(event, func(ctx *etp.Context) error {
		data := new(OwnerTransferRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		if err := services.Control.Admin.TransferGlobalOwnership(
			ctx,
			ctx.Peer.Identity().UserID,
			data.AccountID,
		); err != nil {
			return err
		}

		return socketutils.Respond(ctx, event, struct{}{})
	})
}
