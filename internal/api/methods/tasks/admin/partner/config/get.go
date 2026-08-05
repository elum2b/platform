package config

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Provider    string `json:"provider"     validate:"required"`
	GroupKey    string `json:"group_key"    validate:"required,max=255"`
	Platform    string `json:"platform"     validate:"required"`
}

type GetResponse struct {
	Config tadmin.PartnerConfigModel `json:"config"`
	Found  bool                      `json:"found"`
}

var (
	configGetKey         = "tasks.partner.config.get"
	configGetDescription = `
Returns a partner configuration. Requires the 'tasks.partner.config.get'
permission in the target workspace.`
)

// Get exposes the partner config retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         configGetKey,
	Description: configGetDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(configGetKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		config, found, err := services.Tasks.Admin.GetPartnerConfig(
			ctx.Context,
			data.WorkspaceID,
			data.Provider,
			data.GroupKey,
			data.Platform,
		)

		return GetResponse{Config: config, Found: found}, err
	},
}
