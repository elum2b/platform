package role

import (
	"time"

	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}
type ListResponse struct {
	ID          string    `json:"id"`
	WorkspaceID string    `json:"workspace_id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Position    int32     `json:"position"`
	MemberCount int64     `json:"member_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func List(event string, socket etp.Router) {
	socket.Use(event, middleware.WorkspaceAccess("control.workspace.role.list"))
	socket.On(event, func(ctx *etp.Context) error {
		data := new(ListRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		items, err := services.Control.Admin.ListWorkspaceRoles(
			ctx,
			data.WorkspaceID,
		)
		if err != nil {
			return err
		}

		response := make([]ListResponse, 0, len(items))
		for _, item := range items {
			response = append(
				response,
				ListResponse{
					ID:          item.ID,
					WorkspaceID: item.WorkspaceID,
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
