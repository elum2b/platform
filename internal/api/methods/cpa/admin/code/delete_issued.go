package code

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteIssuedRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CPAID       string `json:"cpa_id"       validate:"required,max=255"`
}
type DeleteIssuedResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteIssuedKey         = "cpa.code.delete_issued"
	deleteIssuedDescription = `
Deletes all issued codes of a CPA offer. Requires the 'cpa.code.delete_issued'
permission in the target workspace.`
)

// DeleteIssued exposes the issued CPA code deletion method.
var DeleteIssued = adapter.Method[DeleteIssuedRequest, DeleteIssuedResponse]{
	Key:         deleteIssuedKey,
	Description: deleteIssuedDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(deleteIssuedKey)},
	Handler: func(ctx *adapter.Context, data DeleteIssuedRequest) (DeleteIssuedResponse, error) {
		affected, err := services.CPA.Admin.DeleteIssuedCodes(
			ctx.Context,
			data.WorkspaceID,
			data.CPAID,
		)

		return DeleteIssuedResponse{Affected: affected}, err
	},
}
