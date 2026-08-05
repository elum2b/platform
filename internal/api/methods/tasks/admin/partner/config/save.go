package config

import (
	"encoding/json"

	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type SaveRequest struct {
	WorkspaceID   string          `json:"workspace_id"        validate:"required,uuid"`
	Provider      string          `json:"provider"            validate:"required"`
	GroupKey      string          `json:"group_key"           validate:"required,max=255"`
	Platform      string          `json:"platform"            validate:"required"`
	IsEnabled     bool            `json:"is_enabled"`
	Secret        *string         `json:"secret,omitempty"`
	WebhookSecret *string         `json:"webhook_secret,omitempty"`
	Target        json.RawMessage `json:"target,omitempty"`
	Settings      json.RawMessage `json:"settings,omitempty"`
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
