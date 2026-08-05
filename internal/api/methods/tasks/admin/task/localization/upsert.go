package localization

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	TaskID      uint64 `json:"task_id"      validate:"required,min=1"`
	Locale      string `json:"locale"       validate:"required,max=32"`
	Title       string `json:"title"        validate:"required"`
	Description string `json:"description"  validate:"required"`
}

var (
	upsertKey         = "tasks.task.localization.upsert"
	upsertDescription = `
Creates or updates a task localization. Requires the
'tasks.task.localization.upsert' permission in the target workspace.`
)

// Upsert exposes the task localization upsert method.
var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(upsertKey),
	},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (struct{}, error) {
		err := services.Tasks.Admin.UpsertTaskLocalization(
			ctx.Context,
			data.WorkspaceID,
			data.TaskID,
			data.Locale,
			data.Title,
			data.Description,
		)

		return struct{}{}, err
	},
}
