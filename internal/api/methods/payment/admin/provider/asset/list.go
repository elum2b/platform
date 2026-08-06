package asset

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	ProviderCode string `json:"provider_code,omitempty"`
	AssetCode    string `json:"asset_code,omitempty"`
	Limit        int32  `json:"limit,omitempty"         validate:"omitempty,min=1,max=100"`
	Offset       int32  `json:"offset,omitempty"        validate:"min=0"`
}

type ListResponse struct {
	ProviderAssets []padm.ProviderAssetModel `json:"provider_assets"`
}

var (
	listKey         = "payment.provider_asset.list"
	listDescription = `
Lists provider assets. Requires the 'payment.provider_asset.list'
permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListProviderAssets(ctx.Context,
			padm.ProviderAssetListParams{
				ProviderCode: d.ProviderCode,
				AssetCode:    d.AssetCode,
				Page:         padm.PageParams{Limit: d.Limit, Offset: d.Offset},
			})
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{ProviderAssets: v}, nil
	},
}
