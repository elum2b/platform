package importapi

import (
	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type PreviewRequest struct {
	WorkspaceID string                 `json:"workspace_id" validate:"required,uuid"`
	Package     refadmin.ExportPackage `json:"package"      validate:"required"`
}

type PreviewResponse struct {
	Preview refadmin.ImportPreview `json:"preview"`
}

var (
	previewKey         = "reference.import.preview"
	previewDescription = `
Validates a reference import package and reports its changes and conflicts.
Requires the 'reference.import.preview' permission in the target workspace.`
)

// Preview exposes the reference import preview method.
var Preview = adapter.Method[PreviewRequest, PreviewResponse]{
	Key:         previewKey,
	Description: previewDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(previewKey),
	},
	Handler: func(ctx *adapter.Context, data PreviewRequest) (PreviewResponse, error) {
		value, err := services.Reference.Admin.PreviewImport(
			ctx.Context,
			data.WorkspaceID,
			data.Package,
		)

		return PreviewResponse{Preview: value}, err
	},
}
