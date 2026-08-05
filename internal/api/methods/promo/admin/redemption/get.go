package redemption

import (
	promoadmin "github.com/elum2b/services/promo/service/admin"
	promouser "github.com/elum2b/services/promo/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID    string `json:"workspace_id"       validate:"required,uuid"`
	ID             uint64 `json:"id"                 validate:"required,min=1"`
	AppID          int64  `json:"app_id"             validate:"required,min=1"`
	PlatformID     int64  `json:"platform_id"        validate:"required,min=1"`
	Platform       string `json:"platform,omitempty"`
	PlatformUserID string `json:"platform_user_id"   validate:"required"`
}

type GetResponse struct {
	Redemption *promoadmin.RedemptionModel `json:"redemption"`
}

var (
	getKey         = "promo.redemption.get"
	getDescription = `
Returns a user's promo redemption. Requires the 'promo.redemption.get'
permission in the target workspace.`
)

// Get exposes the promo redemption retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		value, err := services.Promo.Admin.GetUserRedemption(
			ctx.Context,
			promouser.Identity{
				WorkspaceID:    data.WorkspaceID,
				AppID:          data.AppID,
				PlatformID:     data.PlatformID,
				Platform:       data.Platform,
				PlatformUserID: data.PlatformUserID,
			},
			data.ID,
		)

		return GetResponse{Redemption: value}, err
	},
}
