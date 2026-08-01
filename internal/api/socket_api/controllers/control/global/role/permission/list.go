package permission

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type ListRequest struct {
	RoleID string `json:"role_id" validate:"required,uuid"`
}

func List(event string, socket etp.Router) {
	socket.Use(
		event,
		middleware.GlobalAccess("control.global.role.permission.list"),
	)
	socket.On(event, func(ctx *etp.Context) error {
		data := new(ListRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		items, err := services.Control.Admin.ListGlobalRolePermissions(
			ctx,
			data.RoleID,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(ctx, event, items)
	})
}
