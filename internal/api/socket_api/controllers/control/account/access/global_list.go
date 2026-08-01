package access

import (
	etp "github.com/elum-utils/go-etp"
	internalapi "github.com/elum2b/services/control/service/internalapi"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type AccessGlobalListResponse struct {
	Key      string                  `json:"key"`
	Service  string                  `json:"service"`
	GroupKey string                  `json:"group_key"`
	Scope    internalapi.AccessScope `json:"scope"`
	Position int32                   `json:"position"`
}

func GlobalList(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		methods, err := services.Control.Internal.GetAuthorizedGlobalMethods(
			ctx,
			ctx.Peer.Identity().UserID,
		)
		if err != nil {
			return err
		}

		response := make([]AccessGlobalListResponse, 0, len(methods))
		for _, method := range methods {
			response = append(response, AccessGlobalListResponse{
				Key:      method.Key,
				Service:  method.Service,
				GroupKey: method.GroupKey,
				Scope:    method.Scope,
				Position: method.Position,
			})
		}

		return socketutils.Respond(ctx, event, response)
	})
}
