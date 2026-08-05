package stats

import (
	"time"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type RefreshRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	From        int64  `json:"from"         validate:"required,min=0"`
	Until       int64  `json:"until"        validate:"required,min=0"`
}

var (
	refreshKey         = "payment.daily_stats.refresh"
	refreshDescription = `
Refreshes daily stats. Requires the 'payment.daily_stats.refresh'
permission in the target workspace.`
)

var Refresh = adapter.Method[RefreshRequest, struct{}]{
	Key:         refreshKey,
	Description: refreshDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(refreshKey)},
	Handler: func(ctx *adapter.Context, d RefreshRequest) (struct{}, error) {
		return struct{}{}, services.Payment.Admin.RefreshDailyStats(
			ctx.Context, d.WorkspaceID,
			time.Unix(d.From, 0), time.Unix(d.Until, 0))
	},
}
