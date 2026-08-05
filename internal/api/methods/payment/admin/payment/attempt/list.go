package attempt

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID  string `json:"workspace_id"       validate:"required,uuid"`
	OrderID      uint64 `json:"order_id,omitempty"`
	ProviderCode string `json:"provider_code,omitempty"`
	Status       string `json:"status,omitempty"`
	Limit        int32  `json:"limit,omitempty"    validate:"omitempty,min=1,max=100"`
	Offset       int32  `json:"offset,omitempty"   validate:"min=0"`
}

type ListResponse struct {
	Attempts []padm.PaymentAttemptModel `json:"attempts"`
}

var (
	listKey         = "payment.payment_attempt.list"
	listDescription = `
Lists payment attempts. Requires the 'payment.payment_attempt.list'
permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListPaymentAttempts(ctx.Context,
			padm.AttemptListParams{
				WorkspaceID:  d.WorkspaceID,
				OrderID:      d.OrderID,
				ProviderCode: d.ProviderCode,
				Status:       d.Status,
				Page:         padm.PageParams{Limit: d.Limit, Offset: d.Offset},
			})
		if err != nil {
			return ListResponse{}, err
		}
		return ListResponse{Attempts: v}, nil
	},
}
