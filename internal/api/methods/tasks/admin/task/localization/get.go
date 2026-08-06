package localization

import (
	"github.com/elum2b/services/tasks/repository"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	TaskID      uint64 `json:"task_id"      validate:"required,min=1"`
	Locale      string `json:"locale"       validate:"required,max=32"`
}

type GetResponse struct {
	Localization repository.Localization `json:"localization"`
}

var (
	getKey         = "tasks.task.localization.get"
	getDescription = `
Returns a task localization. Requires the 'tasks.task.localization.get'
permission in the target workspace.`
)

// Get exposes the task localization retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		loc, err := services.Tasks.Admin.GetTaskLocalization(
			ctx.Context,
			data.WorkspaceID,
			data.TaskID,
			data.Locale,
		)

		return GetResponse{Localization: loc}, err
	},
}
