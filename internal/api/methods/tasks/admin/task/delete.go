package task

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "tasks.task.delete"
	deleteDescription = `
Deletes a task. Requires the 'tasks.task.delete' permission in the target
workspace.`
)

// Delete exposes the task deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(deleteKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Tasks.Admin.DeleteTask(
			ctx.Context,
			data.WorkspaceID,
			data.ID,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
