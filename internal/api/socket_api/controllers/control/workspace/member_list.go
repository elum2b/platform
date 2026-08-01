package workspace

import (
	"time"

	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type MemberListRequest struct {
	WorkspaceID string    `json:"workspace_id" validate:"required,uuid"`
	Limit       int32     `json:"limit"        validate:"omitempty,min=1,max=100"`
	CursorAt    time.Time `json:"cursor_at"`
	CursorID    string    `json:"cursor_id"`
}

type MemberListResponse struct {
	WorkspaceID string    `json:"workspace_id"`
	AccountID   string    `json:"account_id"`
	DisplayName string    `json:"display_name"`
	IsOwner     bool      `json:"is_owner"`
	RoleIDs     []string  `json:"role_ids"`
	JoinedAt    time.Time `json:"joined_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func MemberList(event string, socket etp.Router) {
	socket.Use(
		event,
		middleware.WorkspaceAccess("control.workspace.member.list"),
	)

	socket.On(event, func(ctx *etp.Context) error {
		data := new(MemberListRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		members, err := services.Control.Admin.ListMembers(
			ctx,
			data.WorkspaceID,
			admin.Page{
				Limit:    data.Limit,
				CursorAt: data.CursorAt,
				CursorID: data.CursorID,
			},
		)
		if err != nil {
			return err
		}

		response := make([]MemberListResponse, 0, len(members))
		for _, member := range members {
			response = append(response, MemberListResponse{
				WorkspaceID: member.WorkspaceID,
				AccountID:   member.AccountID,
				DisplayName: member.DisplayName,
				IsOwner:     member.IsOwner,
				RoleIDs:     member.RoleIDs,
				JoinedAt:    member.JoinedAt,
				UpdatedAt:   member.UpdatedAt,
			})
		}

		return socketutils.Respond(ctx, event, response)
	})
}
