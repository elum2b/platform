package condition

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID     string `json:"workspace_id"      validate:"required,uuid"`
	ParentTaskID    uint64 `json:"parent_task_id"    validate:"required,min=1"`
	ConditionTaskID uint64 `json:"condition_task_id" validate:"required,min=1"`
}

type GetResponse struct {
	Condition tadmin.ComplexConditionModel `json:"condition"`
}

var (
	getKey         = "tasks.complex_condition.get"
	getDescription = `
Returns a complex task condition. Requires the 'tasks.complex_condition.get'
permission in the target workspace.`
)

// Get exposes the complex condition retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		cond, err := services.Tasks.Admin.GetComplexCondition(
			ctx.Context,
			data.WorkspaceID,
			data.ParentTaskID,
			data.ConditionTaskID,
		)

		return GetResponse{Condition: cond}, err
	},
}
