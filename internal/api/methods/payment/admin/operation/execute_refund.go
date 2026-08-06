package operation

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ExecuteRefundRequest struct {
	WorkspaceID    string `json:"workspace_id"     validate:"required,uuid"`
	OrderID        uint64 `json:"order_id"         validate:"required,min=1"`
	AttemptID      uint64 `json:"attempt_id"       validate:"required,min=1"`
	IdempotencyKey string `json:"idempotency_key"  validate:"required,max=255"`
	AmountMinor    uint64 `json:"amount_minor"     validate:"required,min=0"`
	Reason         string `json:"reason,omitempty"`
}

type ExecuteRefundResponse struct {
	Refund padm.ExecuteRefundResult `json:"refund"`
}

var (
	executeRefundKey         = "payment.operation.execute_refund"
	executeRefundDescription = `
Executes a refund. Requires the 'payment.operation.execute_refund'
permission in the target workspace.`
)

var ExecuteRefund = adapter.Method[ExecuteRefundRequest, ExecuteRefundResponse]{
	Key:         executeRefundKey,
	Description: executeRefundDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(executeRefundKey),
	},
	Handler: func(ctx *adapter.Context, d ExecuteRefundRequest) (ExecuteRefundResponse, error) {
		v, err := services.Payment.Admin.ExecuteRefund(
			ctx.Context,
			padm.ExecuteRefundParams{
				WorkspaceID:    d.WorkspaceID,
				OrderID:        d.OrderID,
				AttemptID:      d.AttemptID,
				IdempotencyKey: d.IdempotencyKey,
				AmountMinor:    d.AmountMinor,
				Reason:         d.Reason,
			})
		if err != nil {
			return ExecuteRefundResponse{}, err
		}

		return ExecuteRefundResponse{Refund: *v}, nil
	},
}
