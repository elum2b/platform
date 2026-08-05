package stats

import (
	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type GetResponse struct {
	Stats refadmin.StatsModel `json:"stats"`
}

var (
	getKey         = "reference.stats.get"
	getDescription = `
Returns aggregate statistics for reference items in a workspace. Requires the
'reference.stats.get' permission in the target workspace.`
)

// Get exposes the reference aggregate statistics method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		value, err := services.Reference.Admin.GetStats(
			ctx.Context,
			data.WorkspaceID,
		)

		return GetResponse{Stats: value}, err
	},
}
