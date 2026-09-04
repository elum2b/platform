package config

import (
	"encoding/json"

	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type SaveRequest struct {
	WorkspaceID   string          `json:"workspace_id"             query:"workspace_id"   validate:"required,uuid"`
	Provider      string          `json:"provider"                 query:"provider"       validate:"required"`
	GroupKey      string          `json:"group_key"                query:"group_key"      validate:"required,max=255"`
	Platform      string          `json:"platform"                 query:"platform"       validate:"required"`
	IsEnabled     bool            `json:"is_enabled"               query:"is_enabled"`
	Secret        *string         `json:"secret,omitempty"         query:"secret"`
	WebhookSecret *string         `json:"webhook_secret,omitempty" query:"webhook_secret"`
	Target        json.RawMessage `json:"target,omitempty"         query:"target"`
	Settings      json.RawMessage `json:"settings,omitempty"       query:"settings"`
}

var (
	configSaveKey         = "tasks.partner.config.save"
	configSaveDescription = `
Creates or updates a partner configuration. Requires the
'tasks.partner.config.save' permission in the target workspace.`
)

// Save exposes the partner config upsert method.
var Save = adapter.Method[SaveRequest, struct{}]{
	Key:         configSaveKey,
	Description: configSaveDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(configSaveKey),
	},
	Handler: func(ctx *adapter.Context, data SaveRequest) (struct{}, error) {
		err := services.Tasks.Admin.SavePartnerConfig(
			ctx.Context,
			tadmin.PartnerConfigModel{
				WorkspaceID:   data.WorkspaceID,
				Provider:      data.Provider,
				GroupKey:      data.GroupKey,
				Platform:      data.Platform,
				IsEnabled:     data.IsEnabled,
				Secret:        data.Secret,
				WebhookSecret: data.WebhookSecret,
				Target:        data.Target,
				Settings:      data.Settings,
			},
		)

		return struct{}{}, err
	},
}
