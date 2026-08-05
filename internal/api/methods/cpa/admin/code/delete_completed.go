package code

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteCompletedRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CPAID       string `json:"cpa_id"       validate:"required,max=255"`
}
type DeleteCompletedResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteCompletedKey         = "cpa.code.delete_completed"
	deleteCompletedDescription = `
Deletes all completed codes of a CPA offer. Requires the
'cpa.code.delete_completed' permission in the target workspace.`
)

// DeleteCompleted exposes the completed CPA code deletion method.
var DeleteCompleted = adapter.Method[DeleteCompletedRequest, DeleteCompletedResponse]{
	Key:         deleteCompletedKey,
	Description: deleteCompletedDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(deleteCompletedKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteCompletedRequest) (DeleteCompletedResponse, error) {
		affected, err := services.CPA.Admin.DeleteCompletedCodes(
			ctx.Context,
			data.WorkspaceID,
			data.CPAID,
		)

		return DeleteCompletedResponse{Affected: affected}, err
	},
}
