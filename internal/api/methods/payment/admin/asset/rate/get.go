package rate

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	AssetCode          string `json:"asset_code"           validate:"required,max=255"`
	ReferenceAssetCode string `json:"reference_asset_code" validate:"required,max=255"`
}

type GetResponse struct {
	Rate padm.AssetRateModel `json:"rate"`
}

var (
	getKey         = "payment.asset_rate.get"
	getDescription = `
Returns an asset rate. Requires the 'payment.asset_rate.get'
permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Payment.Admin.GetAssetRate(
			ctx.Context, d.AssetCode, d.ReferenceAssetCode)

		return GetResponse{Rate: v}, err
	},
}
