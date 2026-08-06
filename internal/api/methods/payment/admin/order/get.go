package order

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
}

type GetResponse struct {
	Order padm.OrderModel `json:"order"`
}

var (
	getKey         = "payment.order.get"
	getDescription = `
Returns an order. Requires the 'payment.order.get'
permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Payment.Admin.GetOrder(
			ctx.Context,
			padm.OrderRefParams{
				WorkspaceID: d.WorkspaceID,
				ID:          d.ID,
			})

		return GetResponse{Order: v}, err
	},
}
