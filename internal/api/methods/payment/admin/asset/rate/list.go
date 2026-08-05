package rate

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	AssetCode          string `json:"asset_code,omitempty"`
	ReferenceAssetCode string `json:"reference_asset_code,omitempty"`
	Limit              int32  `json:"limit,omitempty"            validate:"omitempty,min=1,max=100"`
	Offset             int32  `json:"offset,omitempty"           validate:"min=0"`
}

type ListResponse struct {
	Rates []padm.AssetRateModel `json:"rates"`
}

var (
	listKey         = "payment.asset_rate.list"
	listDescription = `
Lists asset rates. Requires the 'payment.asset_rate.list'
permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListAssetRates(ctx.Context,
			padm.AssetRateListParams{
				AssetCode:          d.AssetCode,
				ReferenceAssetCode: d.ReferenceAssetCode,
				Page:               padm.PageParams{Limit: d.Limit, Offset: d.Offset},
			})
		if err != nil {
			return ListResponse{}, err
		}
		return ListResponse{Rates: v}, nil
	},
}
