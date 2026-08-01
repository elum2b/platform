package method

import (
	"time"

	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type GetRequest struct {
	Key string `json:"key" validate:"required"`
}
type GetResponse struct {
	Key       string            `json:"key"`
	Service   string            `json:"service"`
	GroupKey  string            `json:"group_key"`
	Scope     admin.AccessScope `json:"scope"`
	Position  int32             `json:"position"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func Get(event string, socket etp.Router) {
	socket.Use(event, middleware.GlobalAccess("control.global.method.get"))
	socket.On(event, func(ctx *etp.Context) error {
		data := new(GetRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		item, err := services.Control.Admin.GetMethod(ctx, data.Key)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			GetResponse{
				Key:       item.Key,
				Service:   item.Service,
				GroupKey:  item.GroupKey,
				Scope:     item.Scope,
				Position:  item.Position,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
			},
		)
	})
}
