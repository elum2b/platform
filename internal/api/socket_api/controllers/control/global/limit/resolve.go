package limit

import (
	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type ResolveRequest struct {
	RequestID     string `json:"request_id"     validate:"required,uuid"`
	Approved      bool   `json:"approved"`
	ApprovedLimit int32  `json:"approved_limit" validate:"min=0"`
	Comment       string `json:"comment"        validate:"max=1000"`
}

func Resolve(event string, socket etp.Router) {
	socket.Use(event, middleware.GlobalAccess("control.global.limit.resolve"))
	socket.On(event, func(ctx *etp.Context) error {
		data := new(ResolveRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		item, err := services.Control.Admin.ResolveLimitRequest(
			ctx,
			admin.ResolveLimitRequestParams{
				ActorID:       ctx.Peer.Identity().UserID,
				RequestID:     data.RequestID,
				Approved:      data.Approved,
				ApprovedLimit: data.ApprovedLimit,
				Comment:       data.Comment,
			},
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(ctx, event, item)
	})
}
