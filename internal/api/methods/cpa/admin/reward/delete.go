package reward

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CPAID       string `json:"cpa_id"       validate:"required,max=255"`
	Key         string `json:"key"          validate:"required,max=255"`
}
type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "cpa.reward.delete"
	deleteDescription = `
Deletes a reward from a CPA offer. Requires the 'cpa.reward.delete' permission
in the target workspace.`
)

// Delete exposes the CPA reward deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(deleteKey)},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.CPA.Admin.DeleteReward(
			ctx.Context,
			data.WorkspaceID,
			data.CPAID,
			data.Key,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
