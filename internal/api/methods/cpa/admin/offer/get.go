package offer

import (
	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CPAID       string `json:"cpa_id"       validate:"required,max=255"`
}

type GetResponse struct {
	Offer cpaadmin.OfferModel `json:"offer"`
}

var (
	getKey         = "cpa.offer.get"
	getDescription = `
Returns a CPA offer with its localizations and rewards. Requires the
'cpa.offer.get' permission in the target workspace.`
)

// Get exposes the CPA offer retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		offer, err := services.CPA.Admin.GetOffer(
			ctx.Context,
			data.WorkspaceID,
			data.CPAID,
		)

		return GetResponse{Offer: offer}, err
	},
}
