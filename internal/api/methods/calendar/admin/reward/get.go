package reward

import (
	caluser "github.com/elum2b/services/calendar/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CalendarID  string `json:"calendar_id"  validate:"required,max=255"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
}

type GetResponse struct {
	Reward caluser.RewardModel `json:"reward"`
}

var (
	getKey         = "calendar.reward.get"
	getDescription = `
Returns a single calendar reward. Requires the 'calendar.reward.get'
permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Calendar.Admin.GetReward(
			ctx.Context, d.WorkspaceID, d.CalendarID, d.ID)

		return GetResponse{Reward: v}, err
	},
}
