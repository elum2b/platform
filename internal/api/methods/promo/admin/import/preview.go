package importapi

import (
	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type PreviewRequest struct {
	WorkspaceID string                   `json:"workspace_id" validate:"required,uuid"`
	Package     promoadmin.ExportPackage `json:"package"      validate:"required"`
}

type PreviewResponse struct {
	Preview promoadmin.ImportPreview `json:"preview"`
}

var (
	previewKey         = "promo.import.preview"
	previewDescription = `
Validates a promo import package and reports its changes and conflicts.
Requires the 'promo.import.preview' permission in the target workspace.`
)

// Preview exposes the promo import preview method.
var Preview = adapter.Method[PreviewRequest, PreviewResponse]{
	Key:         previewKey,
	Description: previewDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(previewKey),
	},
	Handler: func(ctx *adapter.Context, data PreviewRequest) (PreviewResponse, error) {
		value, err := services.Promo.Admin.PreviewImport(
			ctx.Context,
			data.WorkspaceID,
			data.Package,
		)

		return PreviewResponse{Preview: value}, err
	},
}
