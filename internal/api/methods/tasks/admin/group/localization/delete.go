package localization

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Key         string `json:"key"          validate:"required,max=255"`
	Locale      string `json:"locale"       validate:"required,max=32"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "tasks.group.localization.delete"
	deleteDescription = `
Deletes a task group localization. Requires the
'tasks.group.localization.delete' permission in the target workspace.`
)

// Delete exposes the group localization deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(deleteKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Tasks.Admin.DeleteGroupLocalization(
			ctx.Context,
			data.WorkspaceID,
			data.Key,
			data.Locale,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
