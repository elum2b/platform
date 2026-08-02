package workspace

import (
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Slug        string `json:"slug"         validate:"required,max=255"`
	Title       string `json:"title"        validate:"required,max=255"`
}

type UpdateResponse struct {
	Affected int64 `json:"affected"`
}

var (
	updateKey         = "control.workspace.update"
	updateDescription = `
Updates the slug and title of a workspace. Requires the
'control.workspace.update' permission in the target workspace.`
)

// Update changes workspace metadata.
var Update = adapter.Method[UpdateRequest, UpdateResponse]{
	Key:         updateKey,
	Description: updateDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.update"),
	},
	Handler: func(ctx *adapter.Context, data UpdateRequest) (UpdateResponse, error) {
		affected, err := services.Control.Admin.UpdateWorkspace(
			ctx.Context,
			controladmin.UpdateWorkspaceParams{
				ActorID: ctx.AccountID, WorkspaceID: data.WorkspaceID,
				Slug: data.Slug, Title: data.Title,
			},
		)
		if err != nil {
			return UpdateResponse{}, err
		}

		return UpdateResponse{Affected: affected}, nil
	},
}
