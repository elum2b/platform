package stats

import (
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ProductID   string `json:"product_id"   validate:"required,max=255"`
	From        int64  `json:"from"         validate:"required,min=0"`
	Until       int64  `json:"until"        validate:"required,min=0"`
}

type ListResponse struct {
	Stats []padm.DailyStatsModel `json:"stats"`
}

var (
	listKey         = "payment.daily_stats.list"
	listDescription = `
Lists daily stats. Requires the 'payment.daily_stats.list'
permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListDailyStats(
			ctx.Context, d.WorkspaceID, d.ProductID,
			time.Unix(d.From, 0), time.Unix(d.Until, 0))
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Stats: v}, nil
	},
}
