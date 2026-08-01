package role

import (
	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type UpdateRequest struct {
	RoleID      string `json:"role_id"     validate:"required,uuid"`
	Title       string `json:"title"       validate:"required,max=255"`
	Description string `json:"description" validate:"max=1000"`
	Position    int32  `json:"position"`
}
type UpdateResponse struct {
	Affected int64 `json:"affected"`
}

func Update(event string, socket etp.Router) {
	socket.Use(event, middleware.GlobalAccess("control.global.role.update"))
	socket.On(event, func(ctx *etp.Context) error {
		data := new(UpdateRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		affected, err := services.Control.Admin.UpdateGlobalRole(
			ctx,
			admin.UpdateRoleParams{
				ActorID:     ctx.Peer.Identity().UserID,
				ID:          data.RoleID,
				Title:       data.Title,
				Description: data.Description,
				Position:    data.Position,
			},
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			UpdateResponse{Affected: affected},
		)
	})
}
