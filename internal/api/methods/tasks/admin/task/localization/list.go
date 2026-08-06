package localization

import (
	"github.com/elum2b/services/tasks/repository"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	TaskID      uint64 `json:"task_id"      validate:"required,min=1"`
}

type ListResponse struct {
	Localizations []repository.Localization `json:"localizations"`
}

var (
	listKey         = "tasks.task.localization.list"
	listDescription = `
Lists localizations for a task. Requires the 'tasks.task.localization.list'
permission in the target workspace.`
)

// List exposes the task localization listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		locs, err := services.Tasks.Admin.ListTaskLocalizations(
			ctx.Context,
			data.WorkspaceID,
			data.TaskID,
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Localizations: locs}, nil
	},
}
