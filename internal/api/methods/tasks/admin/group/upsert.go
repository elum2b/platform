package group

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Key         string `json:"key"          validate:"required,max=255"`
	Position    int32  `json:"position"`
	Active      bool   `json:"active"`
}

var (
	upsertKey         = "tasks.group.upsert"
	upsertDescription = `
Creates or updates a task group. Requires the 'tasks.group.upsert' permission
in the target workspace.`
)

// Upsert exposes the task group upsert method.
var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(upsertKey),
	},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (struct{}, error) {
		err := services.Tasks.Admin.UpsertGroup(
			ctx.Context,
			data.WorkspaceID,
			data.Key,
			data.Position,
			data.Active,
		)

		return struct{}{}, err
	},
}
