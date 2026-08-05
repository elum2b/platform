package reward

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	PromoID     uint64 `json:"promo_id"     validate:"required,min=1"`
	Key         string `json:"key"          validate:"required,max=255"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "promo.reward.delete"
	deleteDescription = `
Deletes a reward from a promo. Requires the 'promo.reward.delete' permission
in the target workspace.`
)

// Delete exposes the promo reward deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(deleteKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Promo.Admin.DeleteReward(
			ctx.Context,
			data.WorkspaceID,
			data.PromoID,
			data.Key,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
