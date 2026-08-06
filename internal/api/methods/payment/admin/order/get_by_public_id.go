package order

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetByPublicIDRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	PublicID    string `json:"public_id"    validate:"required,max=255"`
}

type GetByPublicIDResponse struct {
	Order padm.OrderModel `json:"order"`
}

var (
	getByPublicIDKey         = "payment.order.get_by_public_id"
	getByPublicIDDescription = `
Returns an order by public ID. Requires the 'payment.order.get_by_public_id'
permission in the target workspace.`
)

var GetByPublicID = adapter.Method[GetByPublicIDRequest, GetByPublicIDResponse]{
	Key:         getByPublicIDKey,
	Description: getByPublicIDDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getByPublicIDKey),
	},
	Handler: func(ctx *adapter.Context, d GetByPublicIDRequest) (GetByPublicIDResponse, error) {
		v, err := services.Payment.Admin.GetOrderByPublicID(
			ctx.Context,
			padm.OrderPublicRefParams{
				WorkspaceID: d.WorkspaceID,
				PublicID:    d.PublicID,
			})

		return GetByPublicIDResponse{Order: v}, err
	},
}
