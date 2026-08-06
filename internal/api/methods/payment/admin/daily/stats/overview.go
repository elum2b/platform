package stats

import (
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type OverviewRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	From        int64  `json:"from"         validate:"required,min=0"`
	Until       int64  `json:"until"        validate:"required,min=0"`
}

type OverviewResponse struct {
	Overview []padm.DailyOverviewModel `json:"overview"`
}

var (
	overviewKey         = "payment.daily_stats.overview"
	overviewDescription = `
Lists daily overview. Requires the 'payment.daily_stats.overview'
permission in the target workspace.`
)

var Overview = adapter.Method[OverviewRequest, OverviewResponse]{
	Key:         overviewKey,
	Description: overviewDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(overviewKey)},
	Handler: func(ctx *adapter.Context, d OverviewRequest) (OverviewResponse, error) {
		v, err := services.Payment.Admin.ListDailyOverview(
			ctx.Context, d.WorkspaceID,
			time.Unix(d.From, 0), time.Unix(d.Until, 0))
		if err != nil {
			return OverviewResponse{}, err
		}

		return OverviewResponse{Overview: v}, nil
	},
}
