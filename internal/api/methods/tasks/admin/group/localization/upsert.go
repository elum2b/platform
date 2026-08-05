package localization

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Key         string `json:"key"          validate:"required,max=255"`
	Locale      string `json:"locale"       validate:"required,max=32"`
	Title       string `json:"title"        validate:"required"`
	Description string `json:"description"  validate:"required"`
}

var (
	upsertKey         = "tasks.group.localization.upsert"
	upsertDescription = `
Creates or updates a task group localization. Requires the
'tasks.group.localization.upsert' permission in the target workspace.`
)

// Upsert exposes the group localization upsert method.
var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(upsertKey),
	},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (struct{}, error) {
		err := services.Tasks.Admin.UpsertGroupLocalization(
			ctx.Context,
			data.WorkspaceID,
			data.Key,
			data.Locale,
			data.Title,
			data.Description,
		)

		return struct{}{}, err
	},
}
