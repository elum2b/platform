package member

import (
	"time"

	controlmodel "github.com/elum2b/services/control/model"
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	Limit    int32     `json:"limit,omitempty"     validate:"omitempty,min=1,max=100"`
	CursorAt time.Time `json:"cursor_at,omitempty"`
	CursorID string    `json:"cursor_id,omitempty"`
}

type Item struct {
	AccountID           string                        `json:"account_id"`
	DisplayName         string                        `json:"display_name"`
	Status              controlmodel.MembershipStatus `json:"status"`
	WorkspaceLimit      int32                         `json:"workspace_limit"`
	OwnedWorkspaceCount int64                         `json:"owned_workspace_count"`
	InvitedBy           string                        `json:"invited_by"`
	JoinedAt            time.Time                     `json:"joined_at"`
	UpdatedAt           time.Time                     `json:"updated_at"`
}

type ListResponse struct {
	Members []Item `json:"members"`
}

var (
	listKey         = "control.global.member.list"
	listDescription = `
Lists accounts that are members of the platform. Requires the
'control.global.member.list' global permission.`
)

// List returns global platform members.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.member.list"),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		items, err := services.Control.Admin.ListPlatformMembers(
			ctx.Context,
			controladmin.Page{
				Limit:    data.Limit,
				CursorAt: data.CursorAt,
				CursorID: data.CursorID,
			},
		)
		if err != nil {
			return ListResponse{}, err
		}

		members := make([]Item, 0, len(items))
		for _, item := range items {
			members = append(members, Item{
				AccountID: item.AccountID, DisplayName: item.DisplayName,
				Status: item.Status, WorkspaceLimit: item.WorkspaceLimit,
				OwnedWorkspaceCount: item.OwnedWorkspaceCount,
				InvitedBy:           item.InvitedBy, JoinedAt: item.JoinedAt,
				UpdatedAt: item.UpdatedAt,
			})
		}

		return ListResponse{Members: members}, nil
	},
}
