package code

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteAvailableRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CPAID       string `json:"cpa_id"       validate:"required,max=255"`
}
type DeleteAvailableResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteAvailableKey         = "cpa.code.delete_available"
	deleteAvailableDescription = `
Deletes all available codes of a CPA offer. Requires the
'cpa.code.delete_available' permission in the target workspace.`
)

// DeleteAvailable exposes the available CPA code deletion method.
var DeleteAvailable = adapter.Method[DeleteAvailableRequest, DeleteAvailableResponse]{
	Key:         deleteAvailableKey,
	Description: deleteAvailableDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(deleteAvailableKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteAvailableRequest) (DeleteAvailableResponse, error) {
		affected, err := services.CPA.Admin.DeleteAvailableCodes(
			ctx.Context,
			data.WorkspaceID,
			data.CPAID,
		)

		return DeleteAvailableResponse{Affected: affected}, err
	},
}
