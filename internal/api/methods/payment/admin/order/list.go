package order

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID    string `json:"workspace_id"               validate:"required,uuid"`
	AppID          int64  `json:"app_id,omitempty"`
	PlatformID     int64  `json:"platform_id,omitempty"`
	PlatformUserID string `json:"platform_user_id,omitempty"`
	ProductID      string `json:"product_id,omitempty"`
	Status         string `json:"status,omitempty"`
	Limit          int32  `json:"limit,omitempty"            validate:"omitempty,min=1,max=100"`
	Offset         int32  `json:"offset,omitempty"           validate:"min=0"`
}

type ListResponse struct {
	Orders []padm.OrderModel `json:"orders"`
}

var (
	listKey         = "payment.order.list"
	listDescription = `
Lists orders. Requires the 'payment.order.list'
permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListOrders(ctx.Context,
			padm.OrderListParams{
				WorkspaceID:    d.WorkspaceID,
				AppID:          d.AppID,
				PlatformID:     d.PlatformID,
				PlatformUserID: d.PlatformUserID,
				ProductID:      d.ProductID,
				Status:         d.Status,
				Page: padm.PageParams{
					Limit:  d.Limit,
					Offset: d.Offset,
				},
			})
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Orders: v}, nil
	},
}
