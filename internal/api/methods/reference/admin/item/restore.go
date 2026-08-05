package item

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type RestoreRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Key         string `json:"key"          validate:"required,max=255"`
	Active      bool   `json:"active"`
}

type RestoreResponse struct {
	Affected int64 `json:"affected"`
}

var (
	restoreKey         = "reference.restore"
	restoreDescription = `
Restores a soft-deleted reference item. Requires the 'reference.restore'
permission in the target workspace.`
)

// Restore exposes the reference item restoration method.
var Restore = adapter.Method[RestoreRequest, RestoreResponse]{
	Key:         restoreKey,
	Description: restoreDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(restoreKey),
	},
	Handler: func(ctx *adapter.Context, data RestoreRequest) (RestoreResponse, error) {
		affected, err := services.Reference.Admin.RestoreItem(
			ctx.Context,
			data.WorkspaceID,
			data.Key,
			data.Active,
		)

		return RestoreResponse{Affected: affected}, err
	},
}
