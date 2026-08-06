package price

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type PriceDeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
}

type PriceDeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "payment.price.delete"
	deleteDescription = `
Deletes a price. Requires the 'payment.price.delete'
permission in the target workspace.`
)

var Delete = adapter.Method[PriceDeleteRequest, PriceDeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(deleteKey)},
	Handler: func(ctx *adapter.Context, d PriceDeleteRequest) (PriceDeleteResponse, error) {
		a, err := services.Payment.Admin.DeletePrice(
			ctx.Context, d.WorkspaceID, d.ID)

		return PriceDeleteResponse{Affected: a}, err
	},
}
