package member

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type RemoveRequest struct {
	AccountID string `json:"account_id" validate:"required,uuid"`
}

type RemoveResponse struct {
	Affected int64 `json:"affected"`
}

var (
	removeKey         = "control.global.member.remove"
	removeDescription = `
Removes an account from the platform membership. Requires the
'control.global.member.remove' global permission.`
)

// Remove removes an account from the platform.
var Remove = adapter.Method[RemoveRequest, RemoveResponse]{
	Key:         removeKey,
	Description: removeDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.member.remove"),
	},
	Handler: func(ctx *adapter.Context, data RemoveRequest) (RemoveResponse, error) {
		affected, err := services.Control.Admin.RemovePlatformMember(
			ctx.Context,
			ctx.AccountID,
			data.AccountID,
		)
		if err != nil {
			return RemoveResponse{}, err
		}

		return RemoveResponse{Affected: affected}, nil
	},
}
