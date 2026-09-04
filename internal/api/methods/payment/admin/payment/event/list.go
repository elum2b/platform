package event

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID      string `json:"workspace_id"                query:"workspace_id"      validate:"required,uuid"`
	ProviderCode     string `json:"provider_code,omitempty"     query:"provider_code"`
	ProcessingStatus string `json:"processing_status,omitempty" query:"processing_status"`
	Limit            int32  `json:"limit,omitempty"             query:"limit"             validate:"omitempty,min=1,max=100"`
	Offset           int32  `json:"offset,omitempty"            query:"offset"            validate:"min=0"`
}

type ListResponse struct {
	Events []padm.PaymentEventModel `json:"events"`
}

var (
	listKey         = "payment.payment_event.list"
	listDescription = `
Lists payment events. Requires the 'payment.payment_event.list'
permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListPaymentEvents(ctx.Context,
			padm.EventListParams{
				WorkspaceID:      d.WorkspaceID,
				ProviderCode:     d.ProviderCode,
				ProcessingStatus: d.ProcessingStatus,
				Page: padm.PageParams{
					Limit:  d.Limit,
					Offset: d.Offset,
				},
			})
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Events: v}, nil
	},
}
