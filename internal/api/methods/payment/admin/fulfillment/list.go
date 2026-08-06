package fulfillment

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id"       validate:"required,uuid"`
	OrderID     uint64 `json:"order_id,omitempty"`
	Status      string `json:"status,omitempty"`
	Limit       int32  `json:"limit,omitempty"    validate:"omitempty,min=1,max=100"`
	Offset      int32  `json:"offset,omitempty"   validate:"min=0"`
}

type ListResponse struct {
	Fulfillments []padm.FulfillmentModel `json:"fulfillments"`
}

var (
	listKey         = "payment.fulfillment.list"
	listDescription = `
Lists fulfillments. Requires the 'payment.fulfillment.list'
permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListFulfillments(ctx.Context,
			padm.FulfillmentListParams{
				WorkspaceID: d.WorkspaceID,
				OrderID:     d.OrderID,
				Status:      d.Status,
				Page:        padm.PageParams{Limit: d.Limit, Offset: d.Offset},
			})
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Fulfillments: v}, nil
	},
}
