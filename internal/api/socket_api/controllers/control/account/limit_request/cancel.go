package limitrequest

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type LimitRequestCancelRequest struct {
	RequestID string `json:"request_id" validate:"required,uuid"`
}

type LimitRequestCancelResponse struct {
	Affected int64 `json:"affected"`
}

func Cancel(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		data := new(LimitRequestCancelRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		affected, err := services.Control.Admin.CancelLimitRequest(
			ctx,
			ctx.Peer.Identity().UserID,
			data.RequestID,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			LimitRequestCancelResponse{Affected: affected},
		)
	})
}
