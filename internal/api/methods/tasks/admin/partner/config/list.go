package config

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type ListResponse struct {
	Configs []tadmin.PartnerConfigModel `json:"configs"`
}

var (
	configListKey         = "tasks.partner.config.list"
	configListDescription = `
Lists partner configurations in a workspace. Requires the
'tasks.partner.config.list' permission in the target workspace.`
)

// List exposes the partner config listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         configListKey,
	Description: configListDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(configListKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.Tasks.Admin.ListPartnerConfigs(
			ctx.Context,
			data.WorkspaceID,
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Configs: values}, nil
	},
}
