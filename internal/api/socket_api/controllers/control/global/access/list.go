package access

import (
	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type ListRequest struct {
	Locale string            `json:"locale" validate:"max=32"`
	Scope  admin.AccessScope `json:"scope"  validate:"omitempty,oneof=global workspace"`
}
type ListResponse struct {
	Service     string               `json:"service"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Groups      []admin.AccessGroups `json:"groups"`
}

func List(event string, socket etp.Router) {
	socket.Use(event, middleware.GlobalAccess("control.global.access.list"))
	socket.On(event, func(ctx *etp.Context) error {
		data := new(ListRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		items, err := services.Control.Admin.ListAccess(
			ctx,
			data.Locale,
			data.Scope,
		)
		if err != nil {
			return err
		}

		response := make([]ListResponse, 0, len(items))
		for _, item := range items {
			response = append(
				response,
				ListResponse{
					Service:     item.Service,
					Title:       item.Title,
					Description: item.Description,
					Groups:      item.Groups,
				},
			)
		}

		return socketutils.Respond(ctx, event, response)
	})
}
