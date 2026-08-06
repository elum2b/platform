package cursor

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID  string `json:"workspace_id"            validate:"required,uuid"`
	ProviderCode string `json:"provider_code,omitempty"`
	Network      string `json:"network,omitempty"`
	Limit        int32  `json:"limit,omitempty"         validate:"omitempty,min=1,max=100"`
	Offset       int32  `json:"offset,omitempty"        validate:"min=0"`
}

type ListResponse struct {
	Cursors []padm.ProviderCursorModel `json:"cursors"`
}

var (
	listKey         = "payment.provider_cursor.list"
	listDescription = `
Lists provider cursors. Requires the 'payment.provider_cursor.list'
permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListProviderCursors(ctx.Context,
			padm.ProviderCursorListParams{
				WorkspaceID:  d.WorkspaceID,
				ProviderCode: d.ProviderCode,
				Network:      d.Network,
				Page:         padm.PageParams{Limit: d.Limit, Offset: d.Offset},
			})
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Cursors: v}, nil
	},
}
