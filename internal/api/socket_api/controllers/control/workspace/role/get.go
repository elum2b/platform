package role

import (
	"time"

	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	RoleID      string `json:"role_id"      validate:"required,uuid"`
}
type GetResponse struct {
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

func Get(event string, socket etp.Router) {
	socket.Use(event, middleware.WorkspaceAccess("control.workspace.role.get"))
	socket.On(event, func(ctx *etp.Context) error {
		data := new(GetRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		item, err := services.Control.Admin.GetWorkspaceRole(
			ctx,
			data.WorkspaceID,
			data.RoleID,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			GetResponse{
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
	})
}
