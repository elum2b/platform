package access

import (
	etp "github.com/elum-utils/go-etp"
	internalapi "github.com/elum2b/services/control/service/internalapi"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type AccessWorkspaceListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type AccessWorkspaceListResponse struct {
	Key      string                  `json:"key"`
	Service  string                  `json:"service"`
	GroupKey string                  `json:"group_key"`
	Scope    internalapi.AccessScope `json:"scope"`
	Position int32                   `json:"position"`
}

func WorkspaceList(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		data := new(AccessWorkspaceListRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		methods, err := services.Control.Internal.GetAuthorizedWorkspaceMethods(
			ctx,
			ctx.Peer.Identity().UserID,
			data.WorkspaceID,
		)
		if err != nil {
			return err
		}

		response := make([]AccessWorkspaceListResponse, 0, len(methods))
		for _, method := range methods {
			response = append(response, AccessWorkspaceListResponse{
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
