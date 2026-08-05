package stats

import (
	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
}

type GetResponse struct {
	Stats promoadmin.StatsModel `json:"stats"`
}

var (
	getKey         = "promo.stats.get"
	getDescription = `
Returns aggregate statistics for a promo. Requires the 'promo.stats.get'
permission in the target workspace.`
)

// Get exposes the promo aggregate statistics method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		value, err := services.Promo.Admin.GetStats(
			ctx.Context,
			data.WorkspaceID,
			data.ID,
		)

		return GetResponse{Stats: value}, err
	},
}
