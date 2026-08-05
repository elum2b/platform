package price

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id"     validate:"required,uuid"`
	ProductID   string `json:"product_id"       validate:"required,max=255"`
	AssetCode   string `json:"asset_code,omitempty"`
	Limit       int32  `json:"limit,omitempty"  validate:"omitempty,min=1,max=100"`
	Offset      int32  `json:"offset,omitempty" validate:"min=0"`
}

type ListResponse struct {
	Prices []padm.PriceModel `json:"prices"`
}

var (
	listKey         = "payment.price.list"
	listDescription = `
Lists prices. Requires the 'payment.price.list'
permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListPrices(ctx.Context,
			padm.PriceListParams{
				WorkspaceID: d.WorkspaceID,
				ProductID:   d.ProductID,
				AssetCode:   d.AssetCode,
				Page:        padm.PageParams{Limit: d.Limit, Offset: d.Offset},
			})
		if err != nil {
			return ListResponse{}, err
		}
		return ListResponse{Prices: v}, nil
	},
}
