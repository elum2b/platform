package stats

import (
	"time"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type RefreshRequest struct {
	WorkspaceID string    `json:"workspace_id" validate:"required,uuid"`
	From        time.Time `json:"from"         validate:"required"`
	Until       time.Time `json:"until"        validate:"required"`
}

var (
	refreshKey         = "cpa.stats.refresh"
	refreshDescription = `
Rebuilds daily CPA statistics for a workspace and time range. Requires the
'cpa.stats.refresh' permission in the target workspace.`
)

// Refresh exposes the CPA daily statistics refresh method.
var Refresh = adapter.Method[RefreshRequest, struct{}]{
	Key:         refreshKey,
	Description: refreshDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(refreshKey)},
	Handler: func(ctx *adapter.Context, data RefreshRequest) (struct{}, error) {
		err := services.CPA.Admin.RefreshDailyStats(
			ctx.Context,
			data.WorkspaceID,
			data.From,
			data.Until,
		)

		return struct{}{}, err
	},
}
