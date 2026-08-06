package localization

import (
	"github.com/elum2b/services/tasks/repository"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Key         string `json:"key"          validate:"required,max=255"`
	Locale      string `json:"locale"       validate:"required,max=32"`
}

type GetResponse struct {
	Localization repository.Localization `json:"localization"`
}

var (
	getKey         = "tasks.group.localization.get"
	getDescription = `
Returns a task group localization. Requires the 'tasks.group.localization.get'
permission in the target workspace.`
)

// Get exposes the group localization retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		loc, err := services.Tasks.Admin.GetGroupLocalization(
			ctx.Context,
			data.WorkspaceID,
			data.Key,
			data.Locale,
		)

		return GetResponse{Localization: loc}, err
	},
}
