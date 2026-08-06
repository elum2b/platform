package localization

import (
	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	PromoID     uint64 `json:"promo_id"     validate:"required,min=1"`
	Locale      string `json:"locale"       validate:"required,max=32"`
}

type GetResponse struct {
	Localization promoadmin.LocalizationModel `json:"localization"`
}

var (
	getKey         = "promo.localization.get"
	getDescription = `
Returns a single promo localization. Requires the 'promo.localization.get'
permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Promo.Admin.GetLocalization(
			ctx.Context, d.WorkspaceID, d.PromoID, d.Locale)

		return GetResponse{Localization: v}, err
	},
}
