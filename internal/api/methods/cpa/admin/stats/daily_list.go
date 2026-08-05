package stats

import (
	"time"

	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DailyListRequest struct {
	WorkspaceID string    `json:"workspace_id" validate:"required,uuid"`
	CPAID       string    `json:"cpa_id"       validate:"required,max=255"`
	From        time.Time `json:"from"         validate:"required"`
	Until       time.Time `json:"until"        validate:"required"`
}
type DailyListResponse struct {
	Stats []cpaadmin.DailyStatsModel `json:"stats"`
}

var (
	dailyListKey         = "cpa.stats.daily.list"
	dailyListDescription = `
Lists daily statistics for a CPA offer and time range. Requires the
'cpa.stats.daily.list' permission in the target workspace.`
)

// DailyList exposes the CPA daily statistics listing method.
var DailyList = adapter.Method[DailyListRequest, DailyListResponse]{
	Key:         dailyListKey,
	Description: dailyListDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(dailyListKey)},
	Handler: func(ctx *adapter.Context, data DailyListRequest) (DailyListResponse, error) {
		values, err := services.CPA.Admin.ListDailyStats(
			ctx.Context,
			data.WorkspaceID,
			data.CPAID,
			data.From,
			data.Until,
		)

		return DailyListResponse{Stats: values}, err
	},
}
