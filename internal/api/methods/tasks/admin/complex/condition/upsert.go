package condition

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID     string `json:"workspace_id"      validate:"required,uuid"`
	ParentTaskID    uint64 `json:"parent_task_id"    validate:"required,min=1"`
	ConditionTaskID uint64 `json:"condition_task_id" validate:"required,min=1"`
	RequiredStatus  string `json:"required_status"   validate:"required"`
	Position        int32  `json:"position"`
	IsRequired      bool   `json:"is_required"`
}

var (
	upsertKey         = "tasks.complex_condition.upsert"
	upsertDescription = `
Creates or updates a complex task condition. Requires the
'tasks.complex_condition.upsert' permission in the target workspace.`
)

// Upsert exposes the complex condition upsert method.
var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(upsertKey),
	},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (struct{}, error) {
		err := services.Tasks.Admin.UpsertComplexCondition(
			ctx.Context,
			tadmin.SaveComplexConditionParams{
				WorkspaceID:     data.WorkspaceID,
				ParentTaskID:    data.ParentTaskID,
				ConditionTaskID: data.ConditionTaskID,
				RequiredStatus:  data.RequiredStatus,
				Position:        data.Position,
				IsRequired:      data.IsRequired,
			},
		)

		return struct{}{}, err
	},
}
