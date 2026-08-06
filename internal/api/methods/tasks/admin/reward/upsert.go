package reward

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID string  `json:"workspace_id"   validate:"required,uuid"`
	TaskID      uint64  `json:"task_id"        validate:"required,min=1"`
	Key         string  `json:"key"            validate:"required,max=255"`
	Type        string  `json:"type"           validate:"required,max=255"`
	Quantity    int64   `json:"quantity"`
	Scale       uint16  `json:"scale"`
	Unit        *string `json:"unit,omitempty"`
	Position    int32   `json:"position"       validate:"required,min=1"`
}

var (
	upsertKey         = "tasks.reward.upsert"
	upsertDescription = `
Creates or updates a task reward. Requires the 'tasks.reward.upsert'
permission in the target workspace.`
)

// Upsert exposes the task reward upsert method.
var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(upsertKey),
	},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (struct{}, error) {
		err := services.Tasks.Admin.UpsertReward(
			ctx.Context,
			data.WorkspaceID,
			data.TaskID,
			tadmin.RewardModel{
				Key:      data.Key,
				Type:     data.Type,
				Quantity: data.Quantity,
				Scale:    data.Scale,
				Unit:     data.Unit,
			},
			data.Position,
		)

		return struct{}{}, err
	},
}
