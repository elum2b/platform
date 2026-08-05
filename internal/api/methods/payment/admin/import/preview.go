package importapi

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type PreviewRequest struct {
	WorkspaceID string             `json:"workspace_id" validate:"required,uuid"`
	Package     padm.ExportPackage `json:"package"      validate:"required"`
}

type PreviewResponse struct {
	Preview padm.ImportPreview `json:"preview"`
}

var (
	previewKey         = "payment.import.preview"
	previewDescription = `
Validates a payment import package and reports its changes and conflicts.
Requires the 'payment.import.preview' permission in the target workspace.`
)

var Preview = adapter.Method[PreviewRequest, PreviewResponse]{
	Key:         previewKey,
	Description: previewDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(previewKey),
	},
	Handler: func(ctx *adapter.Context, data PreviewRequest) (PreviewResponse, error) {
		value, err := services.Payment.Admin.PreviewImport(
			ctx.Context,
			data.WorkspaceID,
			data.Package,
		)

		return PreviewResponse{Preview: value}, err
	},
}
