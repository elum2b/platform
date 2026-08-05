package reward

import (
	promouser "github.com/elum2b/services/promo/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	PromoID     uint64 `json:"promo_id"     validate:"required,min=1"`
}

type ListResponse struct {
	Rewards []promouser.RewardModel `json:"rewards"`
}

var (
	listKey         = "promo.reward.list"
	listDescription = `
Lists the rewards of a promo. Requires the 'promo.reward.list' permission in
the target workspace.`
)

// List exposes the promo reward listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.Promo.Admin.ListRewards(
			ctx.Context,
			data.WorkspaceID,
			data.PromoID,
		)

		return ListResponse{Rewards: values}, err
	},
}
