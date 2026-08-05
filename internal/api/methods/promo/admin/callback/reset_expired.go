package callback

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ResetExpiredRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type ResetExpiredResponse struct {
	Affected int64 `json:"affected"`
}

var (
	resetExpiredKey         = "promo.callback.reset_expired"
	resetExpiredDescription = `
Resets expired promo callback processing locks for retry. Requires the
'promo.callback.reset_expired' permission in the target workspace.`
)

// ResetExpired exposes the expired promo callback processing reset method.
var ResetExpired = adapter.Method[ResetExpiredRequest, ResetExpiredResponse]{
	Key:         resetExpiredKey,
	Description: resetExpiredDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(resetExpiredKey),
	},
	Handler: func(ctx *adapter.Context, data ResetExpiredRequest) (ResetExpiredResponse, error) {
		affected, err := services.Promo.Admin.ResetExpiredCallbackProcessing(
			ctx.Context,
			data.WorkspaceID,
		)

		return ResetExpiredResponse{Affected: affected}, err
	},
}
