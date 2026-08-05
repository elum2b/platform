package item

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Key         string `json:"key"          validate:"required,max=255"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "reference.delete"
	deleteDescription = `
Soft-deletes a reference item. Requires the 'reference.delete' permission in
the target workspace.`
)

// Delete exposes the reference item soft-deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(deleteKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Reference.Admin.SoftDeleteItem(
			ctx.Context,
			data.WorkspaceID,
			data.Key,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
