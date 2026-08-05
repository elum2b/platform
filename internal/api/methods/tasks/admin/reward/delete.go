package reward

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	TaskID      uint64 `json:"task_id"      validate:"required,min=1"`
	Key         string `json:"key"          validate:"required,max=255"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "tasks.reward.delete"
	deleteDescription = `
Deletes a task reward. Requires the 'tasks.reward.delete' permission in the
target workspace.`
)

// Delete exposes the task reward deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(deleteKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Tasks.Admin.DeleteReward(
			ctx.Context,
			data.WorkspaceID,
			data.TaskID,
			data.Key,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
