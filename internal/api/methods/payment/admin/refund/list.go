package refund

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id"       query:"workspace_id" validate:"required,uuid"`
	OrderID     uint64 `json:"order_id,omitempty" query:"order_id"`
	Status      string `json:"status,omitempty"   query:"status"`
	Limit       int32  `json:"limit,omitempty"    query:"limit"        validate:"omitempty,min=1,max=100"`
	Offset      int32  `json:"offset,omitempty"   query:"offset"       validate:"min=0"`
}

type ListResponse struct {
	Refunds []padm.RefundModel `json:"refunds"`
}

var (
	listKey         = "payment.refund.list"
	listDescription = `
Lists refunds. Requires the 'payment.refund.list'
permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListRefunds(ctx.Context,
			padm.RefundListParams{
				WorkspaceID: d.WorkspaceID,
				OrderID:     d.OrderID,
				Status:      d.Status,
				Page:        padm.PageParams{Limit: d.Limit, Offset: d.Offset},
			})
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Refunds: v}, nil
	},
}
