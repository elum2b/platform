package localization

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	PromoID     uint64 `json:"promo_id"     validate:"required,min=1"`
	Locale      string `json:"locale"       validate:"required,max=32"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "promo.localization.delete"
	deleteDescription = `
Deletes a promo localization. Requires the 'promo.localization.delete'
permission in the target workspace.`
)

// Delete exposes the promo localization deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(deleteKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Promo.Admin.DeleteLocalization(
			ctx.Context,
			data.WorkspaceID,
			data.PromoID,
			data.Locale,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
