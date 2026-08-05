package localization

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CPAID       string `json:"cpa_id"       validate:"required,max=255"`
	Locale      string `json:"locale"       validate:"required,max=32"`
}
type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "cpa.localization.delete"
	deleteDescription = `
Deletes a CPA offer localization. Requires the 'cpa.localization.delete'
permission in the target workspace.`
)

// Delete exposes the CPA localization deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(deleteKey)},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.CPA.Admin.DeleteLocalization(
			ctx.Context,
			data.WorkspaceID,
			data.CPAID,
			data.Locale,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
