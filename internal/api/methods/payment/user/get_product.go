package user

import (
	puser "github.com/elum2b/services/payment/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetProductRequest struct {
	WorkspaceID string `json:"workspace_id"         validate:"required,uuid"`
	AppID       int64  `json:"app_id"               validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"          validate:"required,min=1"`
	Params      string `json:"params"               validate:"required"`
	ProductID   string `json:"product_id"           validate:"required,max=255"`
	AssetCode   string `json:"asset_code,omitempty"`
	Locale      string `json:"locale,omitempty"`
}

type GetProductResponse struct {
	Product puser.ProductModel `json:"product"`
}

var (
	getProductKey         = "payment.user.get_product"
	getProductDescription = `
Returns a product for the authenticated application user.`
)

var GetProduct = adapter.Method[GetProductRequest, GetProductResponse]{
	Key:         getProductKey,
	Description: getProductDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, d GetProductRequest) (GetProductResponse, error) {
		v, err := services.Payment.User.GetProduct(
			ctx.Context,
			puser.GetProductParams{
				Identity:  *ctx.Identity,
				ProductID: d.ProductID,
				AssetCode: d.AssetCode,
				Locale:    d.Locale,
			},
		)
		if err != nil {
			return GetProductResponse{}, err
		}

		return GetProductResponse{Product: *v}, nil
	},
}
