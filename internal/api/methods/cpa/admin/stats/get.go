package stats

import (
	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CPAID       string `json:"cpa_id"       validate:"required,max=255"`
}
type GetResponse struct {
	Stats cpaadmin.StatsModel `json:"stats"`
}

var (
	getKey         = "cpa.stats.get"
	getDescription = `
Returns aggregate statistics for a CPA offer. Requires the 'cpa.stats.get'
permission in the target workspace.`
)

// Get exposes the CPA aggregate statistics method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		value, err := services.CPA.Admin.GetStats(
			ctx.Context,
			data.WorkspaceID,
			data.CPAID,
		)

		return GetResponse{Stats: value}, err
	},
}
