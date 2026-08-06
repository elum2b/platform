package asset

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	ProviderCode string `json:"provider_code" validate:"required,max=255"`
	AssetCode    string `json:"asset_code"    validate:"required,max=255"`
}

type GetResponse struct {
	ProviderAsset padm.ProviderAssetModel `json:"provider_asset"`
}

var (
	getKey         = "payment.provider_asset.get"
	getDescription = `
Returns a provider asset. Requires the 'payment.provider_asset.get'
permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Payment.Admin.GetProviderAsset(
			ctx.Context, d.ProviderCode, d.AssetCode)

		return GetResponse{ProviderAsset: v}, err
	},
}
