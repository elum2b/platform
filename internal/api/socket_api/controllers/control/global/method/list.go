package method

import (
	"time"

	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type ListResponse struct {
	Key       string            `json:"key"`
	Service   string            `json:"service"`
	GroupKey  string            `json:"group_key"`
	Scope     admin.AccessScope `json:"scope"`
	Position  int32             `json:"position"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func List(event string, socket etp.Router) {
	socket.Use(event, middleware.GlobalAccess("control.global.method.list"))
	socket.On(event, func(ctx *etp.Context) error {
		items, err := services.Control.Admin.ListMethods(ctx)
		if err != nil {
			return err
		}

		response := make([]ListResponse, 0, len(items))
		for _, item := range items {
			response = append(
				response,
				ListResponse{
					Key:       item.Key,
					Service:   item.Service,
					GroupKey:  item.GroupKey,
					Scope:     item.Scope,
					Position:  item.Position,
					CreatedAt: item.CreatedAt,
					UpdatedAt: item.UpdatedAt,
				},
			)
		}

		return socketutils.Respond(ctx, event, response)
	})
}
