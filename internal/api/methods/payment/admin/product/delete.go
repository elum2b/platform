package product

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          string `json:"id"           validate:"required,max=255"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "payment.product.delete"
	deleteDescription = `
Deletes a product. Requires the 'payment.product.delete'
permission in the target workspace.`
)

var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(deleteKey)},
	Handler: func(ctx *adapter.Context, d DeleteRequest) (DeleteResponse, error) {
		a, err := services.Payment.Admin.DeleteProduct(
			ctx.Context, d.WorkspaceID, d.ID)

		return DeleteResponse{Affected: a}, err
	},
}
