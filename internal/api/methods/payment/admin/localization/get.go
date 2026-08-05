package localization

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id"     validate:"required,uuid"`
	Locale      string `json:"locale"           validate:"required,max=255"`
	Key         string `json:"key"              validate:"required,max=255"`
}

type GetResponse struct {
	Localization padm.LocalizationModel `json:"localization"`
}

var (
	getKey         = "payment.localization.get"
	getDescription = `
Returns a localization. Requires the 'payment.localization.get'
permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Payment.Admin.GetLocalization(
			ctx.Context, d.WorkspaceID, d.Locale, d.Key)
		return GetResponse{Localization: v}, err
	},
}
