package reward

import (
	cpauser "github.com/elum2b/services/cpa/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CPAID       string `json:"cpa_id"       validate:"required,max=255"`
}
type ListResponse struct {
	Rewards []cpauser.RewardModel `json:"rewards"`
}

var (
	listKey         = "cpa.reward.list"
	listDescription = `
Lists the rewards of a CPA offer. Requires the 'cpa.reward.list' permission in
the target workspace.`
)

// List exposes the CPA reward listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.CPA.Admin.ListRewards(
			ctx.Context,
			data.WorkspaceID,
			data.CPAID,
		)

		return ListResponse{Rewards: values}, err
	},
}
