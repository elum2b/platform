package workspace

import (
	"time"

	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type MemberListRequest struct {
	WorkspaceID string    `json:"workspace_id"        validate:"required,uuid"`
	Limit       int32     `json:"limit,omitempty"     validate:"omitempty,min=1,max=100"`
	CursorAt    time.Time `json:"cursor_at,omitempty"`
	CursorID    string    `json:"cursor_id,omitempty"`
}

type MemberResponse struct {
	WorkspaceID string    `json:"workspace_id"`
	AccountID   string    `json:"account_id"`
	DisplayName string    `json:"display_name"`
	IsOwner     bool      `json:"is_owner"`
	RoleIDs     []string  `json:"role_ids"`
	JoinedAt    time.Time `json:"joined_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type MemberListResponse struct {
	Members []MemberResponse `json:"members"`
}

var (
	memberListKey         = "control.workspace.member.list"
	memberListDescription = `
Lists members, ownership state, and assigned role IDs in a workspace. Requires
the 'control.workspace.member.list' permission in the target workspace.`
)

// MemberList returns members of a workspace.
var MemberList = adapter.Method[MemberListRequest, MemberListResponse]{
	Key:         memberListKey,
	Description: memberListDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.member.list"),
	},
	Handler: func(
		ctx *adapter.Context,
		data MemberListRequest,
	) (MemberListResponse, error) {
		items, err := services.Control.Admin.ListMembers(
			ctx.Context,
			data.WorkspaceID,
			controladmin.Page{
				Limit:    data.Limit,
				CursorAt: data.CursorAt,
				CursorID: data.CursorID,
			},
		)
		if err != nil {
			return MemberListResponse{}, err
		}

		members := make([]MemberResponse, 0, len(items))
		for _, item := range items {
			members = append(members, MemberResponse{
				WorkspaceID: item.WorkspaceID, AccountID: item.AccountID,
				DisplayName: item.DisplayName, IsOwner: item.IsOwner,
				RoleIDs: item.RoleIDs, JoinedAt: item.JoinedAt,
				UpdatedAt: item.UpdatedAt,
			})
		}

		return MemberListResponse{Members: members}, nil
	},
}
