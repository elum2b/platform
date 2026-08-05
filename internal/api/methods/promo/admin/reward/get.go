package reward

import (
	promouser "github.com/elum2b/services/promo/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	PromoID     uint64 `json:"promo_id"     validate:"required,min=1"`
	Key         string `json:"key"          validate:"required,max=255"`
}

type GetResponse struct {
	Reward promouser.RewardModel `json:"reward"`
}

var (
	getKey         = "promo.reward.get"
	getDescription = `
Returns a single promo reward. Requires the 'promo.reward.get' permission in
the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Promo.Admin.GetReward(
			ctx.Context, d.WorkspaceID, d.PromoID, d.Key)
		return GetResponse{Reward: v}, err
	},
}
