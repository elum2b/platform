package localization

import (
	"github.com/elum2b/services/tasks/repository"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Key         string `json:"key"          validate:"required,max=255"`
}

type ListResponse struct {
	Localizations []repository.Localization `json:"localizations"`
}

var (
	listKey         = "tasks.group.localization.list"
	listDescription = `
Lists localizations for a task group. Requires the
'tasks.group.localization.list' permission in the target workspace.`
)

// List exposes the group localization listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		locs, err := services.Tasks.Admin.ListGroupLocalizations(
			ctx.Context,
			data.WorkspaceID,
			data.Key,
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Localizations: locs}, nil
	},
}
