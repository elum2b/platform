package importapi

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type PreviewRequest struct {
	WorkspaceID string               `json:"workspace_id" validate:"required,uuid"`
	Package     tadmin.ExportPackage `json:"package"      validate:"required"`
}

type PreviewResponse struct {
	Preview tadmin.ImportPreview `json:"preview"`
}

var (
	previewKey         = "tasks.import.preview"
	previewDescription = `
Validates a tasks import package and reports its changes and conflicts.
Requires the 'tasks.import.preview' permission in the target workspace.`
)

// Preview exposes the tasks import preview method.
var Preview = adapter.Method[PreviewRequest, PreviewResponse]{
	Key:         previewKey,
	Description: previewDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(previewKey),
	},
	Handler: func(ctx *adapter.Context, data PreviewRequest) (PreviewResponse, error) {
		value, err := services.Tasks.Admin.PreviewImport(
			ctx.Context,
			data.WorkspaceID,
			data.Package,
		)

		return PreviewResponse{Preview: value}, err
	},
}
