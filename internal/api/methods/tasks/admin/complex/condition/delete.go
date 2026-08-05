package condition

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID     string `json:"workspace_id"      validate:"required,uuid"`
	ParentTaskID    uint64 `json:"parent_task_id"    validate:"required,min=1"`
	ConditionTaskID uint64 `json:"condition_task_id" validate:"required,min=1"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "tasks.complex_condition.delete"
	deleteDescription = `
Deletes a complex task condition. Requires the
'tasks.complex_condition.delete' permission in the target workspace.`
)

// Delete exposes the complex condition deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(deleteKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Tasks.Admin.DeleteComplexCondition(
			ctx.Context,
			data.WorkspaceID,
			data.ParentTaskID,
			data.ConditionTaskID,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
