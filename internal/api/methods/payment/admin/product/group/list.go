package group

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id"     validate:"required,uuid"`
	Limit       int32  `json:"limit,omitempty"  validate:"omitempty,min=1,max=100"`
	Offset      int32  `json:"offset,omitempty" validate:"min=0"`
}

type ListResponse struct {
	Groups []padm.ProductGroupModel `json:"groups"`
}

var (
	listKey         = "payment.product_group.list"
	listDescription = `
Lists product groups in a workspace. Requires the
'payment.product_group.list' permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListProductGroups(ctx.Context,
			padm.ProductGroupListParams{
				WorkspaceID: d.WorkspaceID,
				Page:        padm.PageParams{Limit: d.Limit, Offset: d.Offset},
			})
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Groups: v}, nil
	},
}
