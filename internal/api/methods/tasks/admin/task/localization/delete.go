package localization

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	TaskID      uint64 `json:"task_id"      validate:"required,min=1"`
	Locale      string `json:"locale"       validate:"required,max=32"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "tasks.task.localization.delete"
	deleteDescription = `
Deletes a task localization. Requires the 'tasks.task.localization.delete'
permission in the target workspace.`
)

// Delete exposes the task localization deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(deleteKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Tasks.Admin.DeleteTaskLocalization(
			ctx.Context,
			data.WorkspaceID,
			data.TaskID,
			data.Locale,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
