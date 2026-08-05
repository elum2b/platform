package redemption

import (
	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id"     validate:"required,uuid"`
	ID          uint64 `json:"id"               validate:"required,min=1"`
	Limit       int32  `json:"limit,omitempty"  validate:"omitempty,min=1,max=100"`
	Offset      int32  `json:"offset,omitempty" validate:"min=0"`
}

type ListResponse struct {
	Redemptions []promoadmin.RedemptionModel `json:"redemptions"`
}

var (
	listKey         = "promo.redemption.list"
	listDescription = `
Lists redemptions of a promo. Requires the 'promo.redemption.list' permission
in the target workspace.`
)

// List exposes the promo redemption listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.Promo.Admin.ListRedemptions(
			ctx.Context,
			data.WorkspaceID,
			data.ID,
			promoadmin.Page{Limit: data.Limit, Offset: data.Offset},
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Redemptions: values}, nil
	},
}
