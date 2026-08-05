package importapi

import (
	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type PreviewRequest struct {
	WorkspaceID string                 `json:"workspace_id" validate:"required,uuid"`
	Package     cpaadmin.ExportPackage `json:"package"      validate:"required"`
}

type PreviewResponse struct {
	Preview cpaadmin.ImportPreview `json:"preview"`
}

var (
	previewKey         = "cpa.import.preview"
	previewDescription = `
Validates a CPA import package and reports its changes and conflicts. Requires
the 'cpa.import.preview' permission in the target workspace.`
)

// Preview exposes the CPA import preview method.
var Preview = adapter.Method[PreviewRequest, PreviewResponse]{
	Key:         previewKey,
	Description: previewDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(previewKey)},
	Handler: func(ctx *adapter.Context, data PreviewRequest) (PreviewResponse, error) {
		value, err := services.CPA.Admin.PreviewImport(
			ctx.Context,
			data.WorkspaceID,
			data.Package,
		)

		return PreviewResponse{Preview: value}, err
	},
}
