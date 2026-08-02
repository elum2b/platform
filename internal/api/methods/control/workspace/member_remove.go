package workspace

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type MemberRemoveRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	AccountID   string `json:"account_id"   validate:"required,uuid"`
}

type MemberRemoveResponse struct {
	Affected int64 `json:"affected"`
}

var (
	memberRemoveKey         = "control.workspace.member.remove"
	memberRemoveDescription = `
Removes an account and its role-based access from a workspace. Requires the
'control.workspace.member.remove' permission in the target workspace.`
)

// MemberRemove removes an account from a workspace.
var MemberRemove = adapter.Method[MemberRemoveRequest, MemberRemoveResponse]{
	Key:         memberRemoveKey,
	Description: memberRemoveDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.member.remove"),
	},
	Handler: func(
		ctx *adapter.Context,
		data MemberRemoveRequest,
	) (MemberRemoveResponse, error) {
		affected, err := services.Control.Admin.RemoveMember(
			ctx.Context,
			ctx.AccountID,
			data.WorkspaceID,
			data.AccountID,
		)
		if err != nil {
			return MemberRemoveResponse{}, err
		}

		return MemberRemoveResponse{Affected: affected}, nil
	},
}
