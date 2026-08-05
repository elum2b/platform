package stats

import (
	"time"

	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DailyListRequest struct {
	WorkspaceID string    `json:"workspace_id" validate:"required,uuid"`
	ID          uint64    `json:"id"           validate:"required,min=1"`
	From        time.Time `json:"from"         validate:"required"`
	Until       time.Time `json:"until"        validate:"required"`
}

type DailyListResponse struct {
	Stats []promoadmin.DailyStatsModel `json:"stats"`
}

var (
	dailyListKey         = "promo.stats.daily.list"
	dailyListDescription = `
Lists daily statistics for a promo and time range. Requires the
'promo.stats.daily.list' permission in the target workspace.`
)

// DailyList exposes the promo daily statistics listing method.
var DailyList = adapter.Method[DailyListRequest, DailyListResponse]{
	Key:         dailyListKey,
	Description: dailyListDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(dailyListKey),
	},
	Handler: func(ctx *adapter.Context, data DailyListRequest) (DailyListResponse, error) {
		values, err := services.Promo.Admin.ListDailyStats(
			ctx.Context,
			data.WorkspaceID,
			data.ID,
			data.From,
			data.Until,
		)

		return DailyListResponse{Stats: values}, err
	},
}
