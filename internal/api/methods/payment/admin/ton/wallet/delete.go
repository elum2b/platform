package wallet

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "payment.ton_wallet.delete"
	deleteDescription = `
Deletes a TON wallet configuration. Requires the 'payment.ton_wallet.delete'
permission in the target workspace.`
)

var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(deleteKey)},
	Handler: func(ctx *adapter.Context, d DeleteRequest) (DeleteResponse, error) {
		a, err := services.Payment.Admin.DeleteTONWallet(
			ctx.Context, d.WorkspaceID)
		return DeleteResponse{Affected: a}, err
	},
}
