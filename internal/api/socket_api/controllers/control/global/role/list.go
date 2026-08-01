package role

import (
	"time"

	etp "github.com/elum-utils/go-etp"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type ListResponse struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Position    int32     `json:"position"`
	MemberCount int64     `json:"member_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func List(event string, socket etp.Router) {
	socket.Use(event, middleware.GlobalAccess("control.global.role.list"))
	socket.On(event, func(ctx *etp.Context) error {
		items, err := services.Control.Admin.ListGlobalRoles(ctx)
		if err != nil {
			return err
		}

		response := make([]ListResponse, 0, len(items))
		for _, item := range items {
			response = append(
				response,
				ListResponse{
					ID:          item.ID,
					Code:        item.Code,
					Title:       item.Title,
					Description: item.Description,
					Position:    item.Position,
					MemberCount: item.MemberCount,
					CreatedAt:   item.CreatedAt,
					UpdatedAt:   item.UpdatedAt,
				},
			)
		}

		return socketutils.Respond(ctx, event, response)
	})
}
