package event

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
	Event padm.PaymentEventModel `json:"event"`
}

var (
	getKey         = "payment.payment_event.get"
	getDescription = `
Returns a payment event. Requires the 'payment.payment_event.get'
permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Payment.Admin.GetPaymentEvent(
			ctx.Context,
			padm.EventRefParams{
				WorkspaceID: d.WorkspaceID,
				ID:          d.ID,
			})

		return GetResponse{Event: v}, err
	},
}
