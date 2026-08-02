package workspace

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ArchiveRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type ArchiveResponse struct {
	Affected int64 `json:"affected"`
}

var (
	archiveKey         = "control.workspace.archive"
	archiveDescription = `
Moves the specified workspace to the archived state. Requires the
'control.workspace.archive' permission in the target workspace.`
)

// Archive archives a workspace.
var Archive = adapter.Method[ArchiveRequest, ArchiveResponse]{
	Key:         archiveKey,
	Description: archiveDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.archive"),
	},
	Handler: func(ctx *adapter.Context, data ArchiveRequest) (ArchiveResponse, error) {
		affected, err := services.Control.Admin.ArchiveWorkspace(
			ctx.Context,
			ctx.AccountID,
			data.WorkspaceID,
		)
		if err != nil {
			return ArchiveResponse{}, err
		}

		return ArchiveResponse{Affected: affected}, nil
	},
}
