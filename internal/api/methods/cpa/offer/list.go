package offer

import (
	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id"     validate:"required,uuid"`
	Limit       int32  `json:"limit,omitempty"  validate:"omitempty,min=1,max=100"`
	Offset      int32  `json:"offset,omitempty" validate:"min=0"`
}

type ListResponse struct {
	Offers []cpaadmin.OfferModel `json:"offers"`
}

var (
	listKey         = "cpa.offer.list"
	listDescription = `
Lists CPA offers in a workspace, including localizations and rewards. Requires
the 'cpa.offer.list' permission in the target workspace.`
)

// List exposes the CPA offer listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("cpa.offer.list"),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		offers, err := services.CPA.Admin.ListOffers(
			ctx.Context,
			data.WorkspaceID,
			cpaadmin.Page{Limit: data.Limit, Offset: data.Offset},
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Offers: offers}, nil
	},
}
