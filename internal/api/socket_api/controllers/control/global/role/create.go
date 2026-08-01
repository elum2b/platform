package role

import (
	"time"

	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type CreateRequest struct {
	ID          string `json:"id"          validate:"omitempty,uuid"`
	Code        string `json:"code"        validate:"required,max=255"`
	Title       string `json:"title"       validate:"required,max=255"`
	Description string `json:"description" validate:"max=1000"`
	Position    int32  `json:"position"`
}
type CreateResponse struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Position    int32     `json:"position"`
	MemberCount int64     `json:"member_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func Create(event string, socket etp.Router) {
	socket.Use(event, middleware.GlobalAccess("control.global.role.create"))
	socket.On(event, func(ctx *etp.Context) error {
		data := new(CreateRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		role, err := services.Control.Admin.CreateGlobalRole(
			ctx,
			admin.CreateRoleParams{
				ActorID:     ctx.Peer.Identity().UserID,
				ID:          data.ID,
				Code:        data.Code,
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
			CreateResponse{
				ID:          role.ID,
				Code:        role.Code,
				Title:       role.Title,
				Description: role.Description,
				Position:    role.Position,
				MemberCount: role.MemberCount,
				CreatedAt:   role.CreatedAt,
				UpdatedAt:   role.UpdatedAt,
			},
		)
	})
}
