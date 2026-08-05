package promo

import (
	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id"     validate:"required,uuid"`
	Limit       int32  `json:"limit,omitempty"  validate:"omitempty,min=1,max=100"`
	Offset      int32  `json:"offset,omitempty" validate:"min=0"`
}

type ListResponse struct {
	Promos []promoadmin.PromoModel `json:"promos"`
}

var (
	listKey         = "promo.list"
	listDescription = `
Lists promos in a workspace. Requires the 'promo.list' permission in the
target workspace.`
)

// List exposes the promo listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		promos, err := services.Promo.Admin.ListPromos(
			ctx.Context,
			data.WorkspaceID,
			promoadmin.Page{Limit: data.Limit, Offset: data.Offset},
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Promos: promos}, nil
	},
}
