package localization

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Locale      string `json:"locale"       validate:"required,max=255"`
	Key         string `json:"key"          validate:"required,max=255"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "payment.localization.delete"
	deleteDescription = `
Deletes a localization. Requires the 'payment.localization.delete'
permission in the target workspace.`
)

var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(deleteKey)},
	Handler: func(ctx *adapter.Context, d DeleteRequest) (DeleteResponse, error) {
		a, err := services.Payment.Admin.DeleteLocalization(
			ctx.Context, d.WorkspaceID, d.Locale, d.Key)

		return DeleteResponse{Affected: a}, err
	},
}
