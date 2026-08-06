package refund

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateRequest struct {
	WorkspaceID      string  `json:"workspace_id"                 validate:"required,uuid"`
	OrderID          uint64  `json:"order_id"                     validate:"required,min=1"`
	AttemptID        uint64  `json:"attempt_id"                   validate:"required,min=1"`
	ProviderCode     string  `json:"provider_code"                validate:"required,max=255"`
	ProviderRefundID *string `json:"provider_refund_id,omitempty"`
	AmountMinor      uint64  `json:"amount_minor"                 validate:"required,min=0"`
	AssetCode        string  `json:"asset_code"                   validate:"required,max=255"`
	Status           string  `json:"status,omitempty"`
	Reason           *string `json:"reason,omitempty"`
}

type CreateResponse struct {
	ID uint64 `json:"id"`
}

var (
	createKey         = "payment.refund.create"
	createDescription = `
Creates a refund. Requires the 'payment.refund.create'
permission in the target workspace.`
)

var Create = adapter.Method[CreateRequest, CreateResponse]{
	Key:         createKey,
	Description: createDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(createKey)},
	Handler: func(ctx *adapter.Context, d CreateRequest) (CreateResponse, error) {
		id, err := services.Payment.Admin.CreateRefund(
			ctx.Context,
			padm.RefundCreateParams{
				WorkspaceID:      d.WorkspaceID,
				OrderID:          d.OrderID,
				AttemptID:        d.AttemptID,
				ProviderCode:     d.ProviderCode,
				ProviderRefundID: d.ProviderRefundID,
				AmountMinor:      d.AmountMinor,
				AssetCode:        d.AssetCode,
				Status:           d.Status,
				Reason:           d.Reason,
			})

		return CreateResponse{ID: id}, err
	},
}
