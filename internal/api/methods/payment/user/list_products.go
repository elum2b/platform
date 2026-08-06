package user

import (
	puser "github.com/elum2b/services/payment/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListProductsRequest struct {
	WorkspaceID string `json:"workspace_id"         validate:"required,uuid"`
	AppID       int64  `json:"app_id"               validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"          validate:"required,min=1"`
	Params      string `json:"params"               validate:"required"`
	GroupCode   string `json:"group_code,omitempty"`
	AssetCode   string `json:"asset_code,omitempty"`
	Locale      string `json:"locale,omitempty"`
}

type ListProductsResponse struct {
	Products []puser.ProductModel `json:"products"`
}

var (
	listProductsKey         = "payment.user.list_products"
	listProductsDescription = `
Lists available products for the authenticated application user.`
)

var ListProducts = adapter.Method[ListProductsRequest, ListProductsResponse]{
	Key:         listProductsKey,
	Description: listProductsDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, d ListProductsRequest) (ListProductsResponse, error) {
		v, err := services.Payment.User.ListProducts(
			ctx.Context,
			puser.ListProductsParams{
				Identity:  *ctx.Identity,
				GroupCode: d.GroupCode,
				AssetCode: d.AssetCode,
				Locale:    d.Locale,
			},
		)
		if err != nil {
			return ListProductsResponse{}, err
		}

		return ListProductsResponse{Products: v}, nil
	},
}
