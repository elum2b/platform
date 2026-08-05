package importapi

import (
	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type PreviewRequest struct {
	WorkspaceID string                 `json:"workspace_id" validate:"required,uuid"`
	Package     caladmin.ExportPackage `json:"package"      validate:"required"`
}

type PreviewResponse struct {
	Preview caladmin.ImportPreview `json:"preview"`
}

var (
	previewKey         = "calendar.import.preview"
	previewDescription = `
Validates a calendar import package and reports its changes and conflicts.
Requires the 'calendar.import.preview' permission in the target workspace.`
)

// Preview exposes the calendar import preview method.
var Preview = adapter.Method[PreviewRequest, PreviewResponse]{
	Key:         previewKey,
	Description: previewDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(previewKey),
	},
	Handler: func(ctx *adapter.Context, data PreviewRequest) (PreviewResponse, error) {
		value, err := services.Calendar.Admin.PreviewImport(
			ctx.Context,
			data.WorkspaceID,
			data.Package,
		)

		return PreviewResponse{Preview: value}, err
	},
}
