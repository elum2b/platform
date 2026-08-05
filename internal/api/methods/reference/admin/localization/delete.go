package localization

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ItemKey     string `json:"item_key"     validate:"required,max=255"`
	Locale      string `json:"locale"       validate:"required,max=32"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "reference.localization.delete"
	deleteDescription = `
Deletes a reference item localization. Requires the
'reference.localization.delete' permission in the target workspace.`
)

// Delete exposes the reference localization deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(deleteKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Reference.Admin.DeleteLocalization(
			ctx.Context,
			data.WorkspaceID,
			data.ItemKey,
			data.Locale,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
