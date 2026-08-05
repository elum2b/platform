package product

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          string `json:"id"           validate:"required,max=255"`
}

type GetResponse struct {
	Product padm.ProductModel `json:"product"`
}

var (
	getKey         = "payment.product.get"
	getDescription = `
Returns a product. Requires the 'payment.product.get' permission
in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Payment.Admin.GetProduct(
			ctx.Context, d.WorkspaceID, d.ID)
		return GetResponse{Product: v}, err
	},
}
