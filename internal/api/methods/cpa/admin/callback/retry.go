package callback

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type RetryRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
}
type RetryResponse struct {
	Affected int64 `json:"affected"`
}

var (
	retryKey         = "cpa.callback.retry"
	retryDescription = `
Schedules a CPA callback event for immediate retry. Requires the
'cpa.callback.retry' permission in the target workspace.`
)

// Retry exposes the CPA callback retry method.
var Retry = adapter.Method[RetryRequest, RetryResponse]{
	Key:         retryKey,
	Description: retryDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(retryKey)},
	Handler: func(ctx *adapter.Context, data RetryRequest) (RetryResponse, error) {
		affected, err := services.CPA.Admin.RetryCallbackEventNow(
			ctx.Context,
			data.WorkspaceID,
			data.ID,
		)

		return RetryResponse{Affected: affected}, err
	},
}
