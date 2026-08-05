package transaction

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID  string `json:"workspace_id"            validate:"required,uuid"`
	ProviderCode string `json:"provider_code,omitempty"`
	Network      string `json:"network,omitempty"`
	SourceKey    string `json:"source_key,omitempty"`
	Status       string `json:"status,omitempty"`
	Limit        int32  `json:"limit,omitempty"         validate:"omitempty,min=1,max=100"`
	Offset       int32  `json:"offset,omitempty"        validate:"min=0"`
}

type ListResponse struct {
	Transactions []padm.ProviderTransactionModel `json:"transactions"`
}

var (
	listKey         = "payment.provider_transaction.list"
	listDescription = `
Lists provider transactions. Requires the
'payment.provider_transaction.list' permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListProviderTransactions(ctx.Context,
			padm.ProviderTransactionListParams{
				WorkspaceID:  d.WorkspaceID,
				ProviderCode: d.ProviderCode,
				Network:      d.Network,
				SourceKey:    d.SourceKey,
				Status:       d.Status,
				Page:         padm.PageParams{Limit: d.Limit, Offset: d.Offset},
			})
		if err != nil {
			return ListResponse{}, err
		}
		return ListResponse{Transactions: v}, nil
	},
}
