package item

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id"     validate:"required,uuid"`
	ProductID   string `json:"product_id"       validate:"required,max=255"`
	ItemID      string `json:"item_id,omitempty"`
	Limit       int32  `json:"limit,omitempty"  validate:"omitempty,min=1,max=100"`
	Offset      int32  `json:"offset,omitempty" validate:"min=0"`
}

type ListResponse struct {
	Items []padm.ProductItemModel `json:"items"`
}

var (
	listKey         = "payment.product_item.list"
	listDescription = `
Lists product items. Requires the 'payment.product_item.list'
permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListProductItems(ctx.Context,
			padm.ProductItemListParams{
				WorkspaceID: d.WorkspaceID,
				ProductID:   d.ProductID,
				ItemID:      d.ItemID,
				Page:        padm.PageParams{Limit: d.Limit, Offset: d.Offset},
			})
		if err != nil {
			return ListResponse{}, err
		}
		return ListResponse{Items: v}, nil
	},
}
